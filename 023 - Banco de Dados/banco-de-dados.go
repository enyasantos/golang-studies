package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database := "devbook"
	stringConexao := "root:mysqlpw@tcp(127.0.0.1:49156)/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", stringConexao)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		_, error = db.Exec("CREATE DATABASE " + database)
		if error != nil {
			log.Fatal(error)
		}
	}

	fmt.Println("Conexão está aberta")

	linhas, error := db.Query("SELECT * FROM usuarios")
	if error != nil {
		log.Fatal(error)
	}
	defer linhas.Close()

}
