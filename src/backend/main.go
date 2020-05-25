package main

import (
	"log"
	"net/http"
)

func main() {
	//port := os.Getenv("SERVER_PORT")
	te, err := NewTokenExchange()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", te)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
