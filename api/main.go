package main

import (
	"github.com/kentaro-m/spider/api/driver"
	"github.com/kentaro-m/spider/api/handler"
	"github.com/kentaro-m/spider/api/model"
	"github.com/kentaro-m/spider/api/repository"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/kentaro-m/spider/api/docs"
	"github.com/swaggo/http-swagger"
	"net/http"
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

	articleRepository := repository.NewArticleRepository(connection.SQL)
	articleModel := model.NewArticleModel(articleRepository)
	articleHandler := handler.NewArticleHandler(articleModel)

	r := chi.NewRouter()

	r.Get("/articles", articleHandler.Get)
	r.Post("/articles", articleHandler.Create)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", r)
}
