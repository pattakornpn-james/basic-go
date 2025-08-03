package main

import (
	"basic-go/config"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.LoadConfig()
	psqlDB, err := sqlx.Connect("postgres", cfg.Database().EndPoint())
	if err != nil {
		panic(err)
	}
	_ = psqlDB
	fmt.Println("connect success")
}
