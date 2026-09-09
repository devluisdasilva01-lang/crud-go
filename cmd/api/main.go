package main

import (
	"context"
	"crud-go/infra/database"
	keycloack "crud-go/infra/keycloak"
	"crud-go/internal/cliente"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	context := context.Background()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("Variável de ambiente DATABASE_URL não definida")
	}

	db, err := database.NewPostgresPool(url, context)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()

	keycloakClient, err := keycloak.NewKeycloak(context)
	router := chi.NewRouter()
	cliente.NewRegistrModule(db, router, keycloakClient)

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
