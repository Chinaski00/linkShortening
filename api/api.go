package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"linkShortening/storage"
	"net/http"
)

// RegisterHandlers регистрирует обработчики для методов POST и GET
func RegisterHandlers(router *mux.Router, store storage.Storage) {
	router.HandleFunc("/shorten", shortenURLHandler(store)).Methods("POST")
	router.HandleFunc("/expand/{shortURL}", expandURLHandler(store)).Methods("GET")
}

// shortenURLHandler обрабатывает запрос на сокращение URL
func shortenURLHandler(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Чтение данных из тела запроса
		var input struct {
			URL string `json:"url"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Генерация короткой ссылки и сохранение в хранилище
		shortURL, err := store.ShortenURL(input.URL)
		if err != nil {
			http.Error(w, "Error shortening URL", http.StatusInternalServerError)
			return
		}

		// Возвращение короткой ссылки в ответе
		json.NewEncoder(w).Encode(struct {
			ShortURL string `json:"short_url"`
		}{ShortURL: shortURL})
	}
}

// expandURLHandler обрабатывает запрос на раскрытие короткой ссылки
func expandURLHandler(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлечение переменной из пути
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		// Получение оригинального URL
		originalURL, err := store.ExpandURL(shortURL)
		if err != nil {
			http.Error(w, "Error expanding URL", http.StatusNotFound)
			return
		}

		// Возвращение оригинального URL в ответе
		json.NewEncoder(w).Encode(struct {
			OriginalURL string `json:"original_url"`
		}{OriginalURL: originalURL})
	}
}
