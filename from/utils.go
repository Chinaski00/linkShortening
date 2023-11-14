package from

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const shortURLLength = 10

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateShortURL генерирует случайную короткую ссылку
func GenerateShortURL() string {
	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
