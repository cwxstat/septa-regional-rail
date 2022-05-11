package dbutils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strings"
	"time"
)

func NYtime() time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return time.Now()
	}
	return time.Now().In(loc)

}

func LookupEnv(key string, defaultValue string) string {
	env := defaultValue
	if val, ok := os.LookupEnv(key); ok {
		env = strings.Replace(val, "\n", "", -1)
	}
	return env
}

func Conn(ctx context.Context) (*mongo.Client, error) {

	mongoURI := LookupEnv("MONGO_URI",
		"mongodb://localhost:27017/?directConnection=true&serverSelectionTimeoutMS=2000")

	dbConn, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {

		log.Printf("failed to initialize connection to mongodb: %+v", err)
		return nil, err
	}
	if err := dbConn.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("ping to mongodb failed: %+v", err)
		return nil, err
	}

	return dbConn, nil

}
