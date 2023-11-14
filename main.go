package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"linkShortening/api"
	"linkShortening/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <storage_type>")
		fmt.Println("Valid storage types: 'postgres' or 'in_memory'")
		os.Exit(1)
	}

	router := mux.NewRouter()

	// Выбор хранилища в зависимости от аргумента командной строки
	var store storage.Storage
	switch os.Args[1] {
	case "postgres":
		store, _ = storage.NewPostgresStorage("user=username dbname=shorturlservice sslmode=disable")
	case "in_memory":
		store = storage.NewInMemoryStorage()
	default:
		fmt.Println("Invalid storage option. Use 'postgres' or 'in_memory'.")
		os.Exit(1)
	}

	api.RegisterHandlers(router, store)

	port := "8080" // Порт, на котором будет запущен сервер
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
