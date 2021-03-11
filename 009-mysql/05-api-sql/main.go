package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:carlos123@tcp(127.0.0.1:3306)/kris")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Fprintln(w, "Connected to the MySQL db")
}

func amigos(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:carlos123@tcp(127.0.0.1:3306)/kris")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:carlos123@tcp(127.0.0.1:3306)/kris")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
