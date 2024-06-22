package db

import (
	"database/sql"
	"squaremicroservices/util"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLInterface interface{}

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(config util.Config) (MySQLInterface, error) {
	mySQL := &MySQL{}

	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		return mySQL, err
	}

	mySQL.DB = db

	return mySQL, err
}
