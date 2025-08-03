package main

import (
	"basic-go/config"
	"basic-go/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()
	psqlDB, err := sqlx.Connect("postgres", cfg.Database().EndPoint())
	if err != nil {
		panic(err)
	}

	docs, err := FetchDocument(psqlDB, context.Background())
	if err != nil {
		panic(err)
	}

	for _, doc := range docs {
		fmt.Println("id", doc.Id)
		fmt.Println("content", doc.Content)
	}
}

func FetchDocument(db *sqlx.DB, ctx context.Context) ([]*models.Document, error) {
	var docs = make([]*models.Document, 0)
	query := `
		SELECT
			embedded_documents.id,
			embedded_documents.content
		FROM
			embedded_documents
	`
	if err := db.SelectContext(ctx, &docs, query); err != nil {
		return nil, err
	}

	return docs, nil
}
