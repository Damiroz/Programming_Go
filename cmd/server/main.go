package main

import (
	"log"
	"net/http"

	"example.com/pract6/internal/db"
	"example.com/pract6/internal/httpapi"
	"example.com/pract6/internal/models"
)

func main() {
	d := db.Connect()

	// Автоматическое создание/обновление таблиц (миграция)
	if err := d.AutoMigrate(&models.User{}, &models.Note{}, &models.Tag{}); err != nil {
		log.Fatal("migrate error:", err)
	}

	r := httpapi.BuildRouter(d)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}