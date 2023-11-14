package storage

import "errors"

// ErrURLNotFound представляет ошибку, возникающую при отсутствии URL в хранилище
var ErrURLNotFound = errors.New("URL not found")

// Storage представляет интерфейс для хранилища
type Storage interface {
	ShortenURL(originalURL string) (string, error)
	ExpandURL(shortURL string) (string, error)
}

// NewStorage создает новый экземпляр хранилища в зависимости от переданного типа
func NewStorage(storageType, connectionString string) (Storage, error) {
	switch storageType {
	case "postgres":
		return NewPostgresStorage(connectionString)
	case "in_memory":
		return NewInMemoryStorage(), nil
	default:
		return nil, errors.New("Invalid storage type. Use 'postgres' or 'in_memory'.")
	}
}
