package inmemory

import (
	"errors"
	"testing"
	apperrors "url-shortener/internal/app_errors"
)

func TestInitialization(t *testing.T) {
	s := InitStorage()
	if s.urls == nil {
		t.Fatal("urls storage is nil")
	}
	if len(s.urls) != 0 {
		t.Error("storage len does not equal to zero")
	}
}

func TestAddURLToEmptyStorage(t *testing.T) {
	s := InitStorage()
	alias, err := s.AddURL("url", "alias")
	if err != nil {
		t.Fatal("error does not equal to nil")
	}
	if len(s.urls) != 1 {
		t.Error("storage len does not equal to one")
	}
	if alias != "alias" {
		t.Errorf("incorrect alias in db. expected: \"alias\". got: %v", alias)
	}
}

func TestAddSameURLTwice(t *testing.T) {
	subtests := []struct {
		name   string
		alias  string
		err    error
		length int
	}{
		{"same aliases", "alias", apperrors.ErrAliasAlreadyOccupied, 1},
		{"different aliases", "alias1", apperrors.ErrURLAlreadyExists, 1},
	}

	for _, tt := range subtests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitStorage()
			url := "url"
			alias := "alias"

			s.AddURL(url, alias)
			_, err := s.AddURL(url, tt.alias)
			if !errors.Is(err, tt.err) {
				t.Errorf(
					"incorrect error type: got %v, expected %v",
					err, apperrors.ErrURLAlreadyExists,
				)
			}
			if len(s.urls) != tt.length {
				t.Errorf(
					"storage len must be equal to %v, current len: %v",
					tt.length,
					len(s.urls),
				)
			}
			dburl, _ := s.GetURL(alias)
			if dburl != url {
				t.Errorf("incorrect url in db. expected: %v. got: %v", url, dburl)
			}
		})
	}
}

func TestAddURLToNotEmptyStorage(t *testing.T) {
	s := InitStorage()
	s.AddURL("url1", "alias1")
	_, err := s.AddURL("url2", "alias2")
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	if len(s.urls) != 2 {
		t.Errorf(
			"incorrect storage len, expected: 2, got: %v",
			len(s.urls),
		)
	}
}
