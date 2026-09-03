package main

import (
	"crud-go/infra/database"
	"crud-go/internal/cliente"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("Variável de ambiente DATABASE_URL não definida")
	}

	db, err := database.NewPostgresPool(url)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()

	router := chi.NewRouter()
	cliente.NewRegistrModule(db, router)

	log.Println(
		"Servidor executando em http://localhost:8081",
	)

	err = http.ListenAndServe(
		":8081",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}
