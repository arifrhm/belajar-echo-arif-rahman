package migration

import (
	"github.com/born2ngopi/belajarecho/database"
	"github.com/born2ngopi/belajarecho/internal/model"
)

func Init() {

	dbConn := database.GetConnection()

	var model = []interface{}{
		model.Book{},
	}

	dbConn.AutoMigrate(model...)

}
