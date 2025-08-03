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

	documents, err := FetchDocument(psqlDB, context.Background())
	if err != nil {
		panic(err)
	}

	for _, documents := range documents {
		fmt.Println(documents.DocumentId)
		fmt.Println(documents.Title)
		fmt.Println(documents.FilePath)
		fmt.Println(documents.CategoryId)
		fmt.Println(documents.OwnerId)
		fmt.Println(documents.CategoryId)
		fmt.Println(documents.UpdatedAt)
		fmt.Println(documents.FileSize)
		fmt.Println(documents.FileType)

	}
}

func FetchDocument(db *sqlx.DB, ctx context.Context) ([]*models.Document, error) {
	var docs = make([]*models.Document, 0)
	query := `
		SELECT
			documents.document_id,
			documents.title,
			documents.file_path,
			documents.category_id,
			documents.owner_id,
			documents.created_at,
			documents.updated_at,
			documents.file_size,
			documents.file_type
		FROM
			documents
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

func FetchCategories(db *sqlx.DB, ctx context.Context) ([]*models.Categories, error) {
	var categories = make([]*models.Categories, 0)
	query := `
		SELECT
			categories.category_id,
			categories.category_name,
			categories.description
		FROM
			categories
	`
	err := db.SelectContext(ctx, &categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func Fetchpermissions(db *sqlx.DB, ctx context.Context) ([]*models.Permissions, error) {
	var permissions = make([]*models.Permissions, 0)
	query := `
		SELECT
			permissions.permission_id,
			permissions.document_id,
			permissions.user_id,
			permissions.can_view,
			permissions.can_edit,
			permissions.can_delete,
			permissions.granted_at,
		FROM
			permissions
	`

	err := db.SelectContext(ctx, &permissions, query)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
