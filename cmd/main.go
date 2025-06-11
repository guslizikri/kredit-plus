package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"sigmatech-kredit-plus/internal/router"
	"sigmatech-kredit-plus/pkg"
)

func main() {
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
