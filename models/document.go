package models

type Document struct {
	Id      int64  `db:"id"`
	Content string `db:"content"`
}
