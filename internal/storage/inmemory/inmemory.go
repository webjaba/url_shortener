package inmemory

import (
	apperrors "url-shortener/internal/app_errors"
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
		return "", apperrors.ErrURLNotFound
	}
	return url, nil
}

func (s *MemStorage) AddURL(url, alias string) (string, error) {
	_, exists := s.urls[alias]
	if exists {
		return "", apperrors.ErrAliasAlreadyOccupied
	}
	for existsAlias, existsURL := range s.urls {
		if existsURL == url {
			return existsAlias, apperrors.ErrURLAlreadyExists
		}
	}
	s.urls[alias] = url
	return alias, nil
}
