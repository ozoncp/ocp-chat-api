package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-chat-api/internal/chat_queue"
	"github.com/ozoncp/ocp-chat-api/internal/db"
	"github.com/ozoncp/ocp-chat-api/internal/saver"

	"github.com/rs/zerolog"

	"github.com/ozoncp/ocp-chat-api/internal/chat_service"

	"github.com/ozoncp/ocp-chat-api/pkg/chat_api"

	"github.com/kelseyhightower/envconfig"
	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
	"github.com/ozoncp/ocp-chat-api/internal/chat_repo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var defaultLogger = log.Logger.With().Timestamp().Logger()

func main() {
	if err := Run(); err != nil {
		defaultLogger.Fatal().Err(err).Msg("run application")
	}
}

func Run() error {
	defaultLogger.Info().Msg("Hi, Victor Akhlynin will write this project")

	defaultLogger.Info().Msgf("started service %v", os.Args[0])

	cfg := NewDefaultConfig()
	if err := envconfig.Process("", cfg); err != nil {
		return errors.Wrap(err, "read config from env")
	}

	// persistent storage for chats in Postgres
	ctx := context.Background()
	ctx = defaultLogger.WithContext(ctx)

	dbConfig := &db.DatabaseConf{
		Host:             cfg.DatabaseCfg.Host,
		Port:             cfg.DatabaseCfg.Port,
		User:             cfg.DatabaseCfg.User,
		Password:         cfg.DatabaseCfg.Password,
		DBName:           cfg.DatabaseCfg.DatabaseName,
		Timeout:          defaultSQLTimeout,
		MaxAllowedPacket: defaultSQLMaxAllowedPacket,
		MultiStatements:  defaultSQLMultiStatements,
	}

	migrateConfig := &db.MigrateConf{
		MigrationsURL:      cfg.DatabaseCfg.MigrationsURL,
		MigrationRun:       cfg.DatabaseCfg.MigrationRun,
		MigrationDBVersion: cfg.DatabaseCfg.MigrationDBVersion,
	}

	sqlDB, err := db.InitAndCreateDB(ctx, dbConfig, migrateConfig)
	if err != nil {
		return errors.Wrap(err, "init and create db")
	}

	chatStorage := chat_repo.NewPostgresRepo(sqlDB)

	// our queue Kafka
	kafkaAddr := cfg.KafkaCfg.Host + ":" + cfg.KafkaCfg.Port
	//brokers := []string{kafkaAddr}
	//producer, err := newProducer(brokers)
	//if err != nil {
	//	return errors.Wrap(err, "create kafka producer")
	//}

	consumer, err := sarama.NewConsumer([]string{kafkaAddr}, nil)
	if err != nil {
		return errors.Wrap(err, "new consumer")
	}
	// fixme maybe defer close()

	chatQueue := chat_queue.NewKafkaConsumer(consumer, 4, cfg.KafkaCfg.Topic)

	storageRepoFlusherDeps := chat_flusher.Deps{
		ChunkSize: 2,
	}

	storageRepoFlusher := chat_flusher.NewChatFlusher(storageRepoFlusherDeps)

	storageSaverDeps := &saver.Deps{
		Capacity:    1000,
		FlusherHere: storageRepoFlusher,
		Repository:  chatStorage,
		FlushPeriod: cfg.StorageFlusherPeriod,
		Strategy:    saver.RemoveOldest,
	}
	storageSaver := saver.New(storageSaverDeps)

	serviceDeps := &chat_service.Deps{
		StorageRepo:      chatStorage,
		QueueConsumer:    chatQueue,
		StorageRepoSaver: storageSaver,
	}

	chatService := chat_service.New(serviceDeps)
	chatAPI := chat_api.New(chatService)

	// api
	listener, err := net.Listen(defaultTransportProtocol, cfg.GRPCAddr)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	chat_api.RegisterChatApiServer(grpcServer, chatAPI)

	runner, ctx := errgroup.WithContext(ctx)
	logger := defaultLogger

	runner.Go(func() error {
		return WaitInterruptFromOS(ctx)
	})

	runner.Go(func() error {
		logger.Info().Msg("Start serving grpc api")
		if err := grpcServer.Serve(listener); err != nil {
			return errors.Wrap(err, "grpc server serve")
		}
		return nil
	})

	runner.Go(func() error {
		if err = chatQueue.Run(ctx); err != nil {
			return errors.Wrap(err, "kafka chat queue run")
		}
		return nil
	})

	runner.Go(func() error {
		if err = storageSaver.Run(ctx); err != nil {
			return errors.Wrap(err, "storage saver run")
		}
		return nil
	})

	runner.Go(func() error {
		<-ctx.Done()
		logger.Info().Msg("context done; launching graceful stop of grpc server")
		grpcServer.GracefulStop()
		return nil
	})

	if err := runner.Wait(); err != nil {
		switch {
		case InterruptedFromOS(err):
			defaultLogger.Info().Msg("application stopped by exit signal")
		default:
			return errors.Wrap(err, "runner finished")
		}
	}

	logger.Info().Msg("ChatService stopped")

	return nil
}

var ErrOSSignalInterrupt = errors.New("got interrupt signal from OS")

func WaitInterruptFromOS(ctx context.Context) error {
	logger := *zerolog.Ctx(ctx)
	termChan := make(chan os.Signal, 1)
	signals := DefaultOSSignals()
	signal.Notify(termChan, signals...)
	defer signal.Stop(termChan)

	select {
	case sig := <-termChan:
		logger.Info().Msgf("got signal %v, try graceful shutdown...", sig)
		return ErrOSSignalInterrupt
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "outside interrupt signalwaiting loop")
	}
}

func DefaultOSSignals() []os.Signal {
	return []os.Signal{
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGTERM,
	}
}

func InterruptedFromOS(err error) bool {
	return errors.Is(err, ErrOSSignalInterrupt)
}

func newProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	return producer, err
}

func prepareMessage(topic, mess string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(mess),
	}
	return msg
}
