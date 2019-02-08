package main

import (
	"fmt"
	"log"
	"net/http"
)

func startRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users", process)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Fprintf(w, "Hola mundo"))
}