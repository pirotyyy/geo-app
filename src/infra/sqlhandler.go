package infra

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error)
	}

	sqlHander := new(SqlHandler)
	sqlHander.Conn = conn
	return sqlHander
}
