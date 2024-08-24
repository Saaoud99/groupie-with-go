package main

import (
	"fmt"
	"groupie/Handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handlers.IndexHandler)
	http.HandleFunc("/artists/{id}", Handlers.PageHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("the web server start at port 8080")
}
