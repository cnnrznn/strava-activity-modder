package main

import (
	"log"
	"net/http"
)

func main() {
	//port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/token_exchange", handleTokenExchange)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
