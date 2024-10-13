package main

// @title Wear Store API
// @version 1.0
// @description Wear Store API.
// @host localhost:8080
// @BasePath /api/v1

import (
	"WearStoreAPI/db"
	"WearStoreAPI/internal/handlers"
	"WearStoreAPI/internal/middlewares"
	"WearStoreAPI/pkg/logger"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	logger.LogInit()
	slog.SetDefault(logger.Logger)

	if err := db.DBInit(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	corsMux := middlewares.CorsMiddleware(mux)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: corsMux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	mux.HandleFunc("GET /api/v1/wear/{id}", handlers.GetWearHandler)
	mux.HandleFunc("GET /api/v1/wear", handlers.GetAllWearHandler)
	mux.HandleFunc("POST /api/v1/wear", handlers.PostWearHandler)
	mux.HandleFunc("PATCH /api/v1/wear/{id}", handlers.PatchWearHandler)
	mux.HandleFunc("DELETE /api/v1/wear/{id}", handlers.DeleteWearHandler)

	mux.HandleFunc("/swagger", serveSwagger)

	go func() {
		slog.Info(fmt.Sprintf("Server up with address: %v:%v", host, port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
	}()

	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("Server gracefully stopped")
	}

}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	if p == "" {
		p = "index.html"
	}
	p = filepath.Join("docs", p)
	http.ServeFile(w, r, p)
}
