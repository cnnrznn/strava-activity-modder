package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type Conn struct {
	db *sql.DB
}

func New() *Conn {
	var config map[string]interface{}

	// read db creds
	bytes, err := ioutil.ReadFile("creds.json")
	if err != nil {
		log.Fatal("Couldn't read db credentials", err)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal("Couldn't Unmarshal config", err)
	}

	psqlInfo := fmt.Sprintf(
		"host=%s "+
			"port=%d "+
			"user=%s "+
			"password=%s "+
			"dbname=%s ",
		config["host"], config["port"], config["user"], config["password"],
		config["dbname"],
	)

	// initialize connection pool
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Conn{
		db: db,
	}
}

func (c *Conn) StoreTokens(refresh, access string, expire float64) {
	log.Printf("Stored %v, %v, %v\n", refresh, access, expire)
}
