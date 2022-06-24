package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	internal "github.com/izzxx/username_checker/middleware"
	"github.com/izzxx/username_checker/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("add port in .env file first")
	}

	// Router
	router := chi.NewRouter()

	// middleware
	router.Use(internal.CORS)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	deps := routes.Dependency{
		App: router,
	}

	// Register route endpoint
	deps.Check()

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("server running on port " + server.Addr)
		log.Fatal(server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
