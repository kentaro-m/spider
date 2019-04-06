package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectDB(host, port, userName, password, dbName, dbCharset string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		userName,
		password,
		host,
		port,
		dbName,
		dbCharset,
	)

	db, err := sql.Open("mysql", dbSource)

	if err != nil {
		return nil, err
	}

	dbConn.SQL = db
	return dbConn, err
}
