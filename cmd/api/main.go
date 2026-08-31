package main

import (
	"context"
	"crud-go/internal/cliente"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/cruddb"
	db, err := pgxpool.New(context.Background(), url)

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
