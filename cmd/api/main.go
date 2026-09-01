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

	// Factory
	repository := cliente.NewRepository(db)
	service := cliente.NewService(repository)
	handler := cliente.NewHandler(service)

	router := chi.NewRouter()

	router.Get("/clientes", handler.ListarTodosClientes)
	router.Post("/clientes", handler.AddCliente)
	router.Get("/clientes/{id}", handler.BuscarClientePorId)

	log.Println(
		"Servidor executando em http://localhost:8080",
	)

	err = http.ListenAndServe(
		":8080",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}
