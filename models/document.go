package models

import "time"

type Document struct {
	DocumentId int64     `db:"document_id"`
	Title      string    `db:"title"`
	FilePath   string    `db:"file_path"`
	CategoryId int64     `db:"category_id"`
	OwnerId    int64     `db:"owner_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	FileSize   int64     `db:"file_size"`
	FileType   string    `db:"file_type"`
}
