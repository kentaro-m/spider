package main

import (
	"fmt"
	"github.com/kentaro-m/spider/api/article"
	"github.com/kentaro-m/spider/api/driver"
	"github.com/kentaro-m/spider/api/site"
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

	dbCharset := os.Getenv("DB_CHARSET")
	env := os.Getenv("ENVIRONMENT")

	// Local DB connection info
	dbHost := os.Getenv("LOCAL_DB_HOST")
	dbPort := os.Getenv("LOCAL_DB_PORT")
	dbName := os.Getenv("LOCAL_DB_NAME")
	dbUser := os.Getenv("LOCAL_DB_USER")
	dbPass := os.Getenv("LOCAL_DB_PASSWORD")

	// CloudSQL connection info
	cProjectId := os.Getenv("CLOUDSQL_PROJECT_ID")
	cRegion := os.Getenv("CLOUDSQL_REGION_NAME")
	cInstance := os.Getenv("CLOUDSQL_INSTANCE_NAME")
	cName := os.Getenv("CLOUDSQL_DB_NAME")
	cUser := os.Getenv("CLOUDSQL_USER")
	cPass := os.Getenv("CLOUDSQL_PASSWORD")

	var dsn string

	if env == "prd" {
		dsn = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s:%s:%s)/%s?charset=%s&parseTime=true",
			cUser,
			cPass,
			cProjectId,
			cRegion,
			cInstance,
			cName,
			dbCharset,
		)
	} else {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
			dbCharset,
		)
	}

	connection, err := driver.ConnectDB(dsn)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to connect to DB: %w", err))
		os.Exit(1)
	}

	articleRepository := article.NewArticleRepository(connection)
	siteRepository := site.NewSiteRepository(connection)
	articleModel := article.NewArticleModel(articleRepository, siteRepository)
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
