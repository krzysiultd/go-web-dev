package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// to connect to cloud sql you have to run cloud_sql_proxy and set up account and instance
	db, err := sql.Open("mysql", "root:carlos123@tcp(127.0.0.1:3306)/kris")
	//db, err := sql.Open("mysql", "root:carlos123@cloudsql(strzalkowski:europe-west3:kris-mysql-first)/kris")

	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Succesfully completed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
