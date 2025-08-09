package main

import (
	"cook_book/backend/config"
	"cook_book/backend/internal/db"
	"cook_book/backend/internal/routes"
	"log"
)

func main() {
	db.Connect()
	defer db.Close()

	r := routes.SetupRoutes()

	port := config.InitConfigPort()

	log.Fatal(r.Run(port))
}
