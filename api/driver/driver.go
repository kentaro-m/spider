package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}