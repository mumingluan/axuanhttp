package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/mumingluan/axuanhttp/pkg/config"
	"github.com/mumingluan/axuanhttp/pkg/http"
)

func main() {
	cfg := config.LoadConfig()

	httpsrv := http.NewServer(&cfg)

	if err := httpsrv.LoadSecret(); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpsrv.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	<-done

	{
		log.Info().Msg("HTTP server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()
		if err := httpsrv.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Failed to gracefully shutdown HTTP server")
		}
		log.Info().Msg("HTTP server stopped")
	}
}
