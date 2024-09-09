package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("url", r.RequestURI),
				zap.Duration("duration", duration),
			)
		})
	}
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is route 1"))
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is route 2"))
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is route 3"))
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := chi.NewRouter()

	r.Use(LoggerMiddleware(logger))

	r.Get("/route1", handleRoute1)
	r.Get("/route2", handleRoute2)
	r.Get("/route3", handleRoute3)

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
