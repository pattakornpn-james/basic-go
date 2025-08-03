package models

import "time"

type User struct{
	Id int64 `db:"user_id"`
	UserName string `db:"username"`
	Email string `db:"email"`
	FullName string `db:"full_name"`
	CreatedAt time.Time `db:"created_at"`
}