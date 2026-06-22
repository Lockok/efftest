package main

import (
	"fmt"

	"github.com/Lockok/efftest/internal/config"
	"github.com/Lockok/efftest/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	pool, err := storage.NewPostgres(cfg.DB)
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	fmt.Println("Postgres connection established")
}