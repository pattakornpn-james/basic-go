package models

type Categories struct {
	CategoryId   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
	Description  string `db:"description"`
}
