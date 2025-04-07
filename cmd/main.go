package main

import (
	"net/http"
	"os"
	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"

	"github.com/gorilla/mux"
)

func main() {
	storageType := os.Getenv("STORAGE")

	if storageType == "" {
		panic("storage type must be specified")
	}

	storage := storage.GetStorage(storageType)

	h := handlers.InitHandler(storage)

	r := mux.NewRouter()

	r.HandleFunc("/", h.CreateURL).Methods("POST")
	r.HandleFunc("/{alias}", h.GetURL).Methods("GET")
	http.ListenAndServe(":8888", r)
}
