package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cnnrznn/strava-activity-modder/src/backend/db"
)

type Webhook struct {
	db db.Database
}

func NewWebhook(db db.Database) (*Webhook, error) {
	return &Webhook{
		db: db,
	}, nil
}

func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		wh.handleWebhook(w, r)
	case http.MethodGet:
		wh.handleChallenge(w, r)
	}
}

func (wh *Webhook) handleWebhook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var fields map[string]interface{}
	if err := decoder.Decode(&fields); err != nil {
		log.Println("Error decoding webhook: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(fields)

	if fields["aspect_type"] != "create" ||
		fields["object_type"] != "activity" {
		return
	}

	athleteID := int(fields["owner_id"].(float64))
	activityID := int(fields["object_id"].(float64))

	go wh.renameActivity(athleteID, activityID)
}

func (wh *Webhook) renameActivity(athleteID, activityID int) {
	at, err := wh.db.GetAccessToken(athleteID)
	if err != nil {
		log.Println(err)
		return
	}

	q := url.Values{}
	q.Set("id", strconv.Itoa(activityID))
	q.Set("include_all_efforts", "false")

	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.strava.com/api/v3/activities/%v?", activityID)+q.Encode(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", at))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var activity map[string]interface{}
	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&activity); err != nil {
		log.Println(err)
		return
	}

	movingTime := activity["moving_time"].(float64)
	distance := activity["distance"].(float64)
	averageSpeed := activity["average_speed"].(float64) * 3600 / 1000

	var rate string
	duration := time.Duration(movingTime) * time.Second
	switch activity["type"] {
	case "Run":
		rate = fmt.Sprintf("%.1f min/k", (movingTime/60)/(distance/1000))
	case "Ride":
		wattsP := activity["average_watts"]
		if wattsP != nil {
			rate = fmt.Sprintf("%.1fw", wattsP.(float64))
		} else {
			rate = fmt.Sprintf("%.1fkph", averageSpeed)
		}
	default:
		rate = fmt.Sprintf("%.1fkph", averageSpeed)
	}
	newName := fmt.Sprintf("%v @ %v", duration, rate)

	log.Printf("Renaming %v to %v\n", activity["name"], newName)

	updateAct, err := json.Marshal(map[string]interface{}{
		"commute":     activity["commute"],
		"trainer":     activity["trainer"],
		"description": activity["description"],
		"name":        newName,
		"type":        activity["type"],
		"gear_id":     activity["gear_id"],
	})

	q.Del("include_all_efforts")
	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("https://www.strava.com/api/v3/activities/%v?", activityID)+q.Encode(), bytes.NewReader(updateAct))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", at))
	req.Header.Add("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	log.Println("Put response code: ", resp.Status)

	dec = json.NewDecoder(resp.Body)

	if err := dec.Decode(&activity); err != nil {
		log.Println(err)
		return
	}

	b, err := json.MarshalIndent(activity, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b))
}

func (wh *Webhook) handleChallenge(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("hub.challenge")

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type: ", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(map[string]string{
		"hub.challenge": c,
	})
}
