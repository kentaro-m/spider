package main

import (
	ah "github.com/kentaro-m/spider/api/handler"
	"github.com/kentaro-m/spider/api/driver"
	"log"
	"os"

	"net/http"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
	_ "github.com/kentaro-m/spider/api/docs"
)

// @title Spider API
// @version 1.0
// @description This is a Spider API server.
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbCharset := os.Getenv("DB_CHARSET")

	connection, err := driver.ConnectDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbCharset)

	if err != nil {
		log.Fatal("Error: connecting DB")
	}

	handler := ah.NewArticleHandler(connection)

	r := chi.NewRouter()

	r.Get("/articles", handler.Get)
	r.Post("/articles", handler.Create)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", r)
}
