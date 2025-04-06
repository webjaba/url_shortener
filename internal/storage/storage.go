package storage

import (
	"url-shortener/internal/storage/inmemory"
	"url-shortener/internal/storage/models"
	"url-shortener/internal/storage/postgre"
)

type Storage interface {
	GetURL(alias string) (string, error)
	AddURL(url, alias string) error
}

func GetStorage(storageType string) Storage {
	switch storageType {
	case "postgres":
		db := postgre.ConnectDB()
		db.AutoMigrate(models.Url{})
		return postgre.InitStorage(db)
	case "inmemory":
		return inmemory.InitStorage()
	default:
		return inmemory.InitStorage()
	}
}
