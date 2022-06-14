package seeder

import (
	"log"

	"github.com/born2ngopi/belajarecho/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func BookTableSeeder(conn *gorm.DB) {

	trx := conn.Begin()
	defer trx.Rollback()

	var books = []model.Book{
		{
			Common: model.Common{
				ID: uuid.New(),
			},
			Name: "belajar mvc",
		},
		{
			Common: model.Common{
				ID: uuid.New(),
			},
			Name: "belajar golang",
		},
	}

	if err := trx.Create(&books).Error; err != nil {
		log.Println(err)
		return
	}

	trx.Commit()
}
