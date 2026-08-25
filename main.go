package main

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	cli "crud-go/clientes"
)

func main() {
	url := "postgres://postgres:123456@localhost:5432/cruddb"

	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar", err)
	}

	defer db.Close()

	// cliente := cli.Cliente{
	// 	Nome: "José Pessoa Leal - Dr. Pessoa",
	// 	Email: "drpessoa@gmail.com"
	//  Telefone: "86988556622",
	// }

	// err = cli.CadastrarCliente(db, cliente)

	cliente, err := cli.Ca

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

}
