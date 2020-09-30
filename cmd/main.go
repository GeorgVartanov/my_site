package main

import (
	"github.com/GeorgVartanov/my_site/pkg/databases/postgresDB"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PDB, err := postgresDB.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := PDB.PostgresDB.Close(); err != nil {
			log.Fatal(err)
		}

	}()
}
