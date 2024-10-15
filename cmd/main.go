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
	"WearStoreAPI/internal/repository"
	"WearStoreAPI/internal/service"
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

	productPostgres := repository.ProductRepository{DataBase: db.DB}
	productService := service.NewProductService(&productPostgres)
	productHandlers := handlers.NewProductHandler(productService)

	postWearHandler := http.HandlerFunc(productHandlers.PostWearHandler)
	updateWearHandler := http.HandlerFunc(productHandlers.UpdateWearHandler)
	deleteWearHandler := http.HandlerFunc(productHandlers.DeleteWearHandler)

	mux.HandleFunc("GET /api/v1/wear/{id}", productHandlers.GetWearHandler)
	mux.HandleFunc("GET /api/v1/wear", productHandlers.GetAllWearHandler)
	mux.Handle("POST /api/v1/wear", middlewares.AdminMiddleware(postWearHandler))
	mux.Handle("PUT /api/v1/wear/{id}", middlewares.AdminMiddleware(updateWearHandler))
	mux.Handle("DELETE /api/v1/wear/{id}", middlewares.AdminMiddleware(deleteWearHandler))

	userPostgres := repository.UserRepository{DataBase: db.DB}
	userService := service.NewUserService(&userPostgres)
	userHandlers := handlers.NewUserHandler(userService)

	mux.HandleFunc("GET /api/v1/user/{id}", userHandlers.GetUserHandler)
	mux.HandleFunc("PUT /api/v1/user/{id}", userHandlers.UpdateUserHandler)
	mux.HandleFunc("DELETE /api/v1/user/{id}", userHandlers.DeleteUserHandler)
	mux.HandleFunc("POST /api/v1/login", userHandlers.LoginHandler)
	mux.HandleFunc("POST /api/v1/register", userHandlers.RegisterHandler)

	swaggerHandler := http.HandlerFunc(serveSwagger)
	mux.Handle("/swagger", middlewares.AdminMiddleware(swaggerHandler))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

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
