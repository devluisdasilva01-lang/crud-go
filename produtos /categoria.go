package produtos 

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Categoria struct {
	Id int
	Nome string
}

func addCategoria(db *pgxpool.Pool, categoria Categoria) error {

	sql := db.Exec(
		context.Background(),sql,
		categoria.Nome,
	)

	return err
	
}