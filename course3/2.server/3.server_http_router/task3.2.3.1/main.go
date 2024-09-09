package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func helloWorld2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 2"))
}

func helloWorld3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 3"))
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	r := chi.NewRouter()

	r.Get("/1", helloWorld)
	r.Get("/2", helloWorld2)
	r.Get("/3", helloWorld3)

	log.Printf("Сервер запущен на порту %s", port)

	log.Fatal(http.ListenAndServe(":"+port, r))

}
