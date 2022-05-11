package trainview

import (
	"context"
	"fmt"

	"time"

	"github.com/cwxstat/septa-regional-rail/dbutils"
	"github.com/cwxstat/septa-regional-rail/types"
	"github.com/cwxstat/septa-regional-rail/dbutils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Return struct {
	ID string `json:"_id" bson:"_id"`
	//	Incidents []Incident `json:"incidents" bson:"incidents"`
	Date time.Time `json:"date" bson:"date"`
}

// REF: https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
type Returns struct {
	ID string `json:"_id" bson:"_id"`
	//	Incidents []Incident `json:"incidents" bson:"incidents"`
	Date time.Time `json:"date" bson:"date"`
}

// ActiveSeptaEntry represents the message object returned in the API.
type ActiveSeptaEntry struct {
	MainWebPage string    `json:"mainWebPage" bson:"mainWebPage"`
	TrainView   types.TrainView `json:"trainview" bson:"trainview"`
	Message     string    `json:"message" bson:"message"`
	TimeStamp   time.Time `json:"date" bson:"date"`
}

type trainViewServer struct {
	db db.Database
}

func NewTrainViewServer(ctx context.Context) (*trainViewServer, error) {

	client, err := dbutils.Conn(ctx)
	if err != nil {
		return nil, err
	}

	a := &trainViewServer{
		db: &db.Mongodb{
			Conn:       client,
			Database:   dbutils.LookupEnv("MONGO_DB", "septa"),
			Collection: dbutils.LookupEnv("MONGO_COLLECTION", "events"),
		},
	}
	return a, nil
}

func (a *trainViewServer) Disconnect(ctx context.Context) error {
	return a.db.Disconnect(ctx)
}

func (a *trainViewServer) AddEntry(ctx context.Context, entry *ActiveSeptaEntry) error {
	return a.db.AddEntry(ctx, *entry)
}

func (a *trainViewServer) EntriesMinutesAgo(ctx context.Context, min int) ([]Return, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Only return these fields
	opts := options.Find().SetProjection(bson.D{
		{"incidents", 1},
		{"date", -1},
		{"_id", 1},
	})
	cur, err := a.db.EntriesMinutesAgo(ctx, min, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []Return
	for cur.Next(ctx) {
		var v Return
		if err := cur.Decode(&v); err != nil {
			return nil, fmt.Errorf("decoding mongodb record failed: %+v", err)
		}
		out = append(out, v)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate on mongodb cursor: %+v", err)
	}
	return out, nil

}

func (a *trainViewServer) DatabaseCollection(database string, collection string) {
	a.db.DatabaseCollection(database, collection)
}
