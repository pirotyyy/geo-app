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
	conn, _ := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbPort+")/"+dbName)

	sqlHander := new(SqlHandler)
	sqlHander.Conn = conn
	return sqlHander
}
