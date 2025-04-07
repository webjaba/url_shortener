package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	apperrors "url-shortener/internal/app_errors"
	urlshortener "url-shortener/internal/service/url_shortener"
	"url-shortener/internal/storage"

	"github.com/gorilla/mux"
)

type CreationRequest struct {
	Url string `json:"url"`
}

type CreationResponse struct {
	Alias string `json:"alias"`
}

type Handler struct {
	storage storage.Storage
}

func InitHandler(storage storage.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) sendAlias(alias string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreationResponse{Alias: alias})
}

func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Expected content-type: application/json", http.StatusUnsupportedMediaType)
		return
	}

	req := CreationRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	alias := urlshortener.GenerateRandomAlias()

	newAlias, err := h.storage.AddURL(req.Url, alias)

	if err != nil {
		if errors.Is(err, apperrors.ErrURLAlreadyExists) {
			w.Header().Set("Content-Type", "application/json")
			h.sendAlias(newAlias, w)
			return
		} else if errors.Is(err, apperrors.ErrAliasAlreadyOccupied) {
			for i := 0; i < 10; i++ {
				alias = urlshortener.GenerateRandomAlias()
				newAlias, err = h.storage.AddURL(req.Url, alias)
				if err == nil {
					h.sendAlias(newAlias, w)
					return
				}
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	h.sendAlias(newAlias, w)
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	url, err := h.storage.GetURL(alias)

	if err != nil {
		if errors.Is(err, apperrors.ErrURLNotFound) {
			http.Error(w, "Url not found", http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}
