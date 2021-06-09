package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	for i := 0; i < 5; i++ {
		f, err := os.Open("go.mod")
		if err != nil {
			return errors.Wrap(err, "open file")
		}
		defaultLogger.Info().Msg("open successful")
		defer func() {
			if err := f.Close(); err != nil {
				defaultLogger.Error().Err(err).Msg("close file bad")
			}
		}()
	}

	cfg := NewDefaultConfig()
	if err := envconfig.Process("", cfg); err != nil {
		return errors.Wrap(err, "read config from env")
	}

	// our future persistent DB
	chatStorage := chat_repo.NewRepoInMemory()

	storageFlusherDeps := chat_flusher.Deps{
		ChunkSize: 1,
	}

	// our i/o channel with chat objects
	chatQueue := chat_repo.NewRepoInMemory()

	chatQueueDeps := chat_flusher.Deps{
		ChunkSize: 1,
	}

	chatStorageFlusher := chat_flusher.NewChatFlusher(storageFlusherDeps)
	chatQueueFlusher := chat_flusher.NewChatFlusher(chatQueueDeps)

	serviceDeps := &chat_service.Deps{
		StorageRepo:    chatStorage,
		StorageFlusher: chatStorageFlusher,
		QueueRepo:      chatQueue,
		QueueFlusher:   chatQueueFlusher,
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

	ctx := context.Background()
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

	logger.Info().Msg("Service stopped")

	if err := runner.Wait(); err != nil {
		switch {
		case InterruptedFromOS(err):
			defaultLogger.Info().Msg("application stopped by exit signal")
		default:
			return errors.Wrap(err, "runner finished")
		}
	}
	return nil
}

var ErrOSSignalInterrupt = errors.New("got interrupt signal from OS")

func WaitInterruptFromOS(ctx context.Context) error {
	logger := *zerolog.Ctx(ctx)
	termChan := make(chan os.Signal)
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
