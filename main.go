package main

import (
	"github.com/born2ngopi/belajarecho/database"
	"github.com/born2ngopi/belajarecho/database/migration"
	"github.com/born2ngopi/belajarecho/database/seeder"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	database.Init()

	migration.Init()
	seeder.Seed()

	// seeder
}
