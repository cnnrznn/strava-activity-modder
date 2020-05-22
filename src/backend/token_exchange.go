package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleTokenExchange(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonBs, err := json.MarshalIndent(req.URL.Query(), "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonBs))
}
