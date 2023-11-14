package storage

import (
	"linkShortening/from"
	"sync"
)

// InMemoryStorage представляет in-memory хранилище
type InMemoryStorage struct {
	urls  map[string]string
	mutex sync.RWMutex
}

// NewInMemoryStorage создает новый экземпляр in-memory хранилища
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]string),
	}
}

// ShortenURL сохраняет оригинальный URL и возвращает сокращенный
func (s *InMemoryStorage) ShortenURL(originalURL string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	shortURL := from.GenerateShortURL()

	// Сохранение соответствия короткой и оригинальной ссылок
	s.urls[shortURL] = originalURL

	return shortURL, nil
}

// ExpandURL возвращает оригинальный URL по короткой ссылке
func (s *InMemoryStorage) ExpandURL(shortURL string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Поиск оригинального URL по короткой ссылке
	originalURL, found := s.urls[shortURL]
	if !found {
		return "", ErrURLNotFound
	}

	return originalURL, nil
}
