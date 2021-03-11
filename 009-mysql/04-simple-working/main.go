package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:carlos123@tcp(127.0.0.1:3306)/kris")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert, err := db.Query("INSERT INTO users VALUES('ELLIOTT')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// fmt.Println("Succesfully inserted into the db")

	rows, err := db.Query("SELECT name FROM amigos")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		s += name + "\n"
	}
	fmt.Println(s)
}
