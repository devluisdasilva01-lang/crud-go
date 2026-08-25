package clientes

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"	
)

type Cliente struct {
	id int 
	Nome string 
	Email string
	Telefone string 
}

// Funcionalidades 
func CadastrarCliente(db *pgxpool.Pool, cliente Cliente) error {

	sql := `
		INSERT INTO clientes (nome, email, telefone)
			VALUES ($1, $2, $3)
		RETURNING id
	`
	_, err := db.Exec(
		context.Background(), sql,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
	)

	return err 
}

