package db

import (
	"database/sql"
	"squaremicroservices/util"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB      *sql.DB
	QUERIES *Queries
}

func NewMySQL(config util.Config) (*MySQL, error) {
	mySQL := &MySQL{}

	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		return mySQL, err
	}

	mySQL.DB = db
	mySQL.QUERIES = New(db)

	return mySQL, err
}
