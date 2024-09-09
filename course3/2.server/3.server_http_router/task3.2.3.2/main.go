package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func helloGroup(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	num := chi.URLParam(r, "num")
	w.Write([]byte("Group " + name + " Привет, мир " + num))
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	r := chi.NewRouter()

	r.Get("/group{name}/{num}", helloGroup)

	log.Printf("Сервер запущен на порту %s", port)

	log.Fatal(http.ListenAndServe(":"+port, r))

}
