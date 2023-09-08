package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Suitmedia!")
}

func main() {
	log.Println("Starting web server")

	http.HandleFunc("/", IndexHandler)

	log.Println("Server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
