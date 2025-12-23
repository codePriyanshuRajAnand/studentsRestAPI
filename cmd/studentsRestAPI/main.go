package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/config"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/http/handlers/students"
	"github.com/codePriyanshuRajAnand/studentsRestAPI/internal/storage/sqlite"
)

func main() {

	// get Config

	cfg := config.MustLoad()

	// Database initialization

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized successfully", slog.String("env", cfg.Env), slog.String("version", "v1.0.0"))
	// setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", students.Create(storage))

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Starting Server", slog.String("Address", cfg.HttpAddress.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("Internal Server Error %s", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error: ", err.Error()))
	}

	slog.Info("Server shutdown successfully!")
}
