package seeder

import "github.com/born2ngopi/belajarecho/database"

func Seed() {

	dbConn := database.GetConnection()

	BookTableSeeder(dbConn)

}
