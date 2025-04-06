package main

import (
	"os"
	"url-shortener/internal/storage"
)

func main() {
	storageType := os.Getenv("STORAGE")

	if storageType == "" {
		panic("storage type must be specified")
	}

	storage := storage.GetStorage(storageType)

	_ = storage

	// TODO: init url handler
}
