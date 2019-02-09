package main

import (
	"fmt"
	"log"
	"net/http"
)

func startRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/users", users)
	http.HandleFunc("/users/", solo)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Fprintf(w, "<h1>Sdneirf API</h1>"))
}