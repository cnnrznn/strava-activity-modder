package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	conn *mongo.Client
}

func New() *MongoDB {
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

	return &MongoDB{
		conn: conn,
	}
}

func (db *MongoDB) StoreTokens(refresh, access string, expires float64) error {
	log.Println(refresh, access, expires)
	return nil
}
