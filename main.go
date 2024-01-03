package main

import (
	"chat/database"
	"chat/database/migrations"
	"chat/src"
	"log"
	"net/http"
)

func main() {
	database.Connection()
	go migrations.Migrate()

	src.HandleRouting()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
