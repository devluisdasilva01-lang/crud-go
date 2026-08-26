package main

import (
	"context"
	"fmt"
	//"log"
	"github.com/jackc/pgx/v5/pgxpool"
	//cli "crud-go/clientes"
	p "crud-go/produtos"
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

	//cliente, err := cli.CarregarClientePeloId(db, 1)
	//if err != nil {
	//	log.Fatal("Erro ao conectar", err)
	//}

	//categoria := p.Categoria {
	//	Nome: "Telecomunicação",
	//}

	//err = p.AddCategoria(db, categoria)

	categoria, err := p.CarregarTodasCategorias(db)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	//fmt.Printf("%d - %s - %s - %s\n", cliente.Id, cliente.Nome, cliente.Email, cliente.Telefone)

	// for _, cliente := range clientes  {
	// 	fmt.Println(
	// 		cliente.Id,
	// 		cliente.Nome,
	// 		cliente.Email,
	// 		cliente.Telefone,
	// 	)

	for _, categoria := range categorias {
		fmt.Printf("%d - %s\n", categoria.Id, categoria.Nome)
	}
	
}
