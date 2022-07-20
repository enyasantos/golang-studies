package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o MySQL
)

func Connection() (*sql.DB, error) {
	database := "devbook"
	stringConnection := "root:mysqlpw@tcp(127.0.0.1:49156)/" + database + "?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
