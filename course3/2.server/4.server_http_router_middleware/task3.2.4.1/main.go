package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rr, r)

		duration := time.Since(start)
		log.Printf("method=%s url=%s status=%d duration=%s", r.Method, r.RequestURI, rr.statusCode, duration)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
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

	r := chi.NewRouter()

	r.Use(LoggingMiddleware)

	r.Get("/route1", handleRoute1)
	r.Get("/route2", handleRoute2)
	r.Get("/route3", handleRoute3)

	log.Printf("Сервер запущен на порту %s", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
