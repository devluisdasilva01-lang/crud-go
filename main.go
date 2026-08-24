package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	url := "postgres://postgres:123456@localhost:5432/cruddb"

	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		panic("Erro ao conectar")
	}

	defer db.Close()

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

}
