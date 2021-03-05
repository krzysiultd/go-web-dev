package main

import "net/http"

func init() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}

func main() {

}
