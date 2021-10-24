package main

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/server"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/database"
	"log"
)

func main() {
	database.Setup()
	router := server.Setup()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
