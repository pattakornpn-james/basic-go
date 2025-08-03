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
	
	abc, err := FetchUsers(psqlDB, context.Background())
	if err != nil {
		panic(err)
	}
	for _, item := range abc {
		fmt.Println(item)
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

func FetchUsers(db *sqlx.DB, ctx context.Context) ([]*models.User, error) {
	var users = make([]*models.User, 0)
	query := `
		SELECT
			users.user_id,
			users.username,
			users.email,
			users.full_name,
			users.created_at
		FROM
			users
	`
	err := db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
