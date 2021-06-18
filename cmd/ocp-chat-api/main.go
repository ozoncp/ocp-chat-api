package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ozoncp/ocp-chat-api/internal/db"
	"github.com/ozoncp/ocp-chat-api/internal/saver"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/ozoncp/ocp-chat-api/internal/chat_service"

	"github.com/ozoncp/ocp-chat-api/pkg/chat_api"

	"github.com/jackc/pgx"
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

	// persistent storage interaction
	ctx := context.Background()

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

	// our i/o channel with chat objects
	chatQueue := chat_repo.NewRepoInMemory() // future Kafka

	// statistics module
	statisticsRepo := chat_repo.NewRepoInMemory()

	statisticsFlusherDeps := chat_flusher.Deps{
		ChunkSize: 1,
	}

	statisticsFlusher := chat_flusher.NewChatFlusher(statisticsFlusherDeps)

	statisticSaverDeps := &saver.Deps{
		Capacity:    1000,
		FlusherHere: statisticsFlusher,
		Repository:  statisticsRepo,
		FlushPeriod: 10 * time.Second,
		Strategy:    saver.RemoveOldest,
	}
	statisticsSaver := saver.New(statisticSaverDeps)

	serviceDeps := &chat_service.Deps{
		StorageRepo:     chatStorage,
		QueueRepo:       chatQueue,
		StatisticsSaver: statisticsSaver,
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

// PostgresFormatDSN  format: postgres://username:password@localhost:5432/database_name
func PostgresFormatDSN(c *pgx.ConnConfig) string {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
	return psqlInfo
}

func NewPostgresSQL(ctx context.Context, psqlConfig *pgx.ConnConfig) (*sql.DB, error) {
	logger := *zerolog.Ctx(ctx)
	logger.Info().Str("component", "database").Str("stage", "create").Msgf("opening connection")

	sqlDB, err := sql.Open("postgres", PostgresFormatDSN(psqlConfig))
	if err != nil {
		return nil, errors.Wrap(err, "open database")
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping database")
	}

	return sqlDB, nil
}
