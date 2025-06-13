package main

import (
	"log"

	"github.com/go-playground/validator/v10"

	_ "github.com/joho/godotenv/autoload"

	"sigmatech-kredit-plus/internal/router"
	"sigmatech-kredit-plus/pkg"
)

var Validate *validator.Validate

func main() {
	pkg.InitValidator()

	db, err := pkg.Posql()
	if err != nil {
		log.Fatal("ini error db start", err)
	}

	router := router.NewRouter(db)
	server := pkg.Server(router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
