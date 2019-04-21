package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(host, port, userName, password, dbName, dbCharset string) (*sql.DB, error) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		userName,
		password,
		host,
		port,
		dbName,
		dbCharset,
	)

	return sql.Open("mysql", dbSource)
}
