package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	conn *mongo.Client
}

func New() *Mongodb {
	uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	conn, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Mongodb{
		conn: conn,
	}
}

func (db *Mongodb) StoreTokens(refresh, access string, expires float64) error {
	log.Println(refresh, access, expires)
	return nil
}
