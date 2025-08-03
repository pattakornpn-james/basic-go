package main

import (
	"basic-go/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Host", cfg.Database().Host())
}
