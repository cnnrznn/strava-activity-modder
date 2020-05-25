package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/cnnrznn/strava-activity-modder/src/backend/db"
	"github.com/cnnrznn/strava-activity-modder/src/backend/db/memory"
)

const (
	wantScope = "read,activity:write,activity:read"
	clientID  = "48402"
	urlToken  = "https://www.strava.com/api/v3/oauth/token"
)

type TokenExchange struct {
	ClientSecret string
	db           db.Database
}

func NewTokenExchange() (*TokenExchange, error) {
	// read client secret file
	bytes, err := ioutil.ReadFile("client_secret.txt")
	if err != nil {
		log.Println("Failed to read client secret file")
		return nil, err
	}

	return &TokenExchange{
		ClientSecret: string(bytes),
		db:           memory.New(),
	}, nil
}

func (te *TokenExchange) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	code := values.Get("code")
	scope := values.Get("scope")

	if scope != wantScope {
		http.Error(w, "We need more permissions", http.StatusBadRequest)
		return
	}

	// GET request for tokens
	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("client_secret", te.ClientSecret)
	q.Set("code", code)
	q.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(urlToken, q)
	if err != nil {
		http.Error(w, "Problem exchanging code with Strava", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Could not parse token exchange body", http.StatusInternalServerError)
		return
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(body, &obj); err != nil {
		http.Error(w, "Failed to parse access tokens", http.StatusInternalServerError)
		return
	}

	id := int(obj["athlete"].(map[string]interface{})["id"].(float64))
	rt := obj["refresh_token"].(string)
	at := obj["access_token"].(string)
	ea := int(obj["expires_at"].(float64))

	te.db.StoreTokens(rt, at, id, ea)

	pobj, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(pobj)
}
