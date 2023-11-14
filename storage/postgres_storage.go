package storage

import (
	"database/sql"
	"linkShortening/from"
	"log"

	_ "github.com/lib/pq"
)

// PostgresStorage представляет хранилище для PostgreSQL
type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage создает новый экземпляр хранилища для PostgreSQL
func NewPostgresStorage(connectionString string) (*PostgresStorage, error) {
	// Используйте параметры подключения к вашей базе данных PostgreSQL
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

// ShortenURL сохраняет оригинальный URL и возвращает сокращенный
func (s *PostgresStorage) ShortenURL(originalURL string) (string, error) {
	// Генерация уникальной короткой ссылки
	shortURL := from.GenerateShortURL()

	// Пример запроса INSERT
	query := "INSERT INTO short_urls (short_url, original_url) VALUES ($1, $2) RETURNING short_url"
	var insertedShortURL string
	err := s.db.QueryRow(query, shortURL, originalURL).Scan(&insertedShortURL)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return insertedShortURL, nil
}

// ExpandURL возвращает оригинальный URL по короткой ссылке
func (s *PostgresStorage) ExpandURL(shortURL string) (string, error) {
	// Пример запроса SELECT
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM short_urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrURLNotFound
		}
		log.Println(err)
		return "", err
	}

	return originalURL, nil
}
