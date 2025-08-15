package db

import (
	"cook_book/backend/config"
	"cook_book/backend/internal/model"
	"cook_book/backend/internal/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := config.InitConfigDsn()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.SendLog("DB [Connect]", "gorm.Open", err)
		panic(err)
	}

	if err = DB.AutoMigrate(&model.CookBook{}, &model.User{}); err != nil {
		utils.SendLog("DB [Connect]", "DB.AutoMigrate", err)
		panic(err)
	}
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		utils.SendLog("DB [Close]", "DB.DB", err)
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		utils.SendLog("DB [Close]", "db.Close", err)
		log.Fatal(err)
	}
}

func Get() *gorm.DB {
	return DB
}
