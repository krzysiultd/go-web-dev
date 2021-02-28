package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("KS-Key", "this is from Chris")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Anything you want</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
