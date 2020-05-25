package main

import (
	"net/http"

	"github.com/cnnrznn/strava-activity-modder/src/backend/db"
)

func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
	case http.MethodGet:
	}
}

type Webhook struct {
	db db.Database
}

func NewWebhook(db db.Database) (*Webhook, error) {
	return &Webhook{
		db: db,
	}, nil
}
