package main

import (
	"log"
	"net/http"
	"net/url"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	params := url.Values{}
	params.Set("client_id", "48402")
	params.Set("redirect_uri", "http://activitymodder.com")
	params.Set("response_type", "code")
	params.Set("scope", "profile:read_all,activity:write")
	q := params.Encode()
	http.Redirect(w, req, "https://strava.com/oauth/authorize?"+q, http.StatusPermanentRedirect)
}

func main() {
	http.HandleFunc("/", redirect)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
