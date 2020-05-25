package main

import (
	"log"
	"net/http"

	"github.com/cnnrznn/strava-activity-modder/src/backend/db/memory"
)

func main() {
	//port := os.Getenv("SERVER_PORT")
	db := memory.New()

	te, err := NewTokenExchange(db)
	if err != nil {
		log.Fatal(err)
	}
	wh, err := NewWebhook(db)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", te)
	http.Handle("/webhook", wh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
