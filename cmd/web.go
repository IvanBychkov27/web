package main

import (
	"context"
	"fmt"
	"github.com/ivanbychkov27/web/internal/application"
	"github.com/ivanbychkov27/web/internal/config"
	"go.uber.org/zap"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	version = "undefined"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("error create logger, %v", err)
		os.Exit(1)
	}

	logger.Info("web", zap.String("version", version))

	err = run(logger)
	if err != nil {
		logger.Error("error run application", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("done")
}

func run(logger *zap.Logger) error {
	cfg := config.New()
	err := cfg.Load()
	if err != nil {
		return fmt.Errorf("error load config, %w", err)
	}

	logger.Info("config loaded", zap.Any("config", cfg))

	ln, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return fmt.Errorf("error create main listener, %w", err)
	}
	defer ln.Close()

	app := application.New(logger, cfg)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go app.Run(cancel, wg, ln)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	signal.Notify(signals, syscall.SIGINT)

	select {
	case <-signals:
		logger.Info("terminate by signal")
		cancel()
	case <-ctx.Done():
		logger.Info("terminate by context")
	}
	app.Stop()

	wg.Wait()

	return nil
}
