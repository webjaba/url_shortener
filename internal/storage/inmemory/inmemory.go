package inmemory

import (
	"errors"
)

type MemStorage struct {
	urls map[string]string
}

func InitStorage() *MemStorage {
	return &MemStorage{urls: make(map[string]string)}
}

func (s *MemStorage) GetURL(alias string) (string, error) {
	url, exists := s.urls[alias]
	if !exists {
		return "", errors.New("url does not exists")
	}
	return url, nil
}

func (s *MemStorage) AddURL(url, alias string) error {
	_, exists := s.urls[alias]
	if exists {
		return errors.New("url exists")
	}
	return nil
}
