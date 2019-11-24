package main

import (
	"database/sql"
	"fmt"

	//mySQL drive
	_ "github.com/go-sql-driver/mysql"
)

func createCon() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/menu_db")
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())

	}
	return db
}
