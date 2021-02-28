package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy style")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty cat")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c) // exact match, doesn't run the code for /cat/something/else like dog handler

	http.ListenAndServe(":8080", nil)
}
