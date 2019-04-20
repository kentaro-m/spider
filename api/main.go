package main

import (
	"github.com/kentaro-m/spider/api/article"
	"github.com/kentaro-m/spider/api/driver"
	"golang.org/x/xerrors"
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
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to load .env file: %w", err))
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbCharset := os.Getenv("DB_CHARSET")

	connection, err := driver.ConnectDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbCharset)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to connect to DB: %w", err))
		os.Exit(1)
	}

	articleRepository := article.NewArticleRepository(connection)
	articleModel := article.NewArticleModel(articleRepository)
	articleHandler := article.NewArticleHandler(articleModel)

	r := chi.NewRouter()

	r.Get("/articles", articleHandler.Get)
	r.Post("/articles", articleHandler.Create)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	err = http.ListenAndServe(":8080", r)

	if err != nil {
		log.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}
