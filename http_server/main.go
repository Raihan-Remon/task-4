package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello welcome to index page</h1><br><a href="/about">About</a>`)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello welcome to about page</h1><br><a href="/index">Index</a>`)
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server started at http://localhost:8080/index")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
