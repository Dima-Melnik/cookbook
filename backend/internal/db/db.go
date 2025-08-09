package db

import (
	"cook_book/backend/config"
	"cook_book/backend/internal/model"
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
		panic(err)
	}

	if err = DB.AutoMigrate(&model.CookBook{}, &model.User{}); err != nil {
		panic(err)
	}
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func Get() *gorm.DB {
	return DB
}
