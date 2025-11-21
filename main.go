package main

import (
	"crud-api/pkg/models"
	"crud-api/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connStr := os.Getenv("connStr")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.Book{})
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r, db)
	log.Fatal(http.ListenAndServe(":8080", r))
	//connStr=postgresql://neondb_owner:npg_tSGrbOs6g2VX@ep-delicate-morning-a46fcge7-pooler.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require

}
