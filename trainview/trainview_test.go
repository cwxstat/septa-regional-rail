package trainview

import (
	"context"

	"fmt"
	"testing"
	"time"

	"github.com/cwxstat/septa-regional-rail/dbutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFull(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	as, err := NewTrainViewServer(ctx)
	if err != nil {
		t.FailNow()
	}

	as.DatabaseCollection("test", "test")

	iwebp := []IncidentWebPage{
		IncidentWebPage{
			Page: "Page1",
		},
		IncidentWebPage{
			Page: "Page2",
		},
	}

	incident1 := Incident{
		IncidentNo:      "1",
		IncidentType:    "test",
		IncidentSubTupe: "",
		Location:        "",
		Municipality:    "",
		DispatchTime:    "",
		Station:         "",
		IncidentStatus:  []IncidentStatus{},
	}

	incident2 := Incident{
		IncidentNo:      "2",
		IncidentType:    "test",
		IncidentSubTupe: "",
		Location:        "",
		Municipality:    "",
		DispatchTime:    "",
		Station:         "",
		IncidentStatus:  []IncidentStatus{},
	}

	err = as.db.AddEntry(ctx, ActiveSeptaEntry{
		MainWebPage:      "Main",
		IncidentWebPages: iwebp,
		Incidents:        []Incident{incident1, incident2},
		Message:          "Test Message",
		TimeStamp:        dbutils.NYtime(),
	})
	if err != nil {
		t.FailNow()
	}
	opts := options.Find().SetProjection(bson.D{
		{"incidents", 1},
		{"date", -1},
		{"_id", 1},
	})
	cur, err := as.db.EntriesMinutesAgo(ctx, 1, opts)
	if err != nil {
		t.FailNow()
	}
	defer cur.Close(ctx)
	var out []Return
	for cur.Next(ctx) {
		var v Return
		if err := cur.Decode(&v); err != nil {
			t.Error(err)

		}
		out = append(out, v)
	}
	if err := cur.Err(); err != nil {
		t.Error(err)
	}

	if out[0].Incidents[1].IncidentNo != "2" {
		t.Fatalf("Didn't get correct value back. Expected 2, got: %+v\n", out[1].Incidents[1].IncidentNo)
	}

	err = as.db.DeleteAll(ctx, "Test Message")
	if err != nil {
		t.FailNow()
	}

	err = as.db.Disconnect(ctx)
	if err != nil {
		t.FailNow()
	}

}

func TestConn(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Client
		wantErr bool
	}{
		{
			name: "Simple connection test",
			args: args{
				ctx: context.TODO(),
			},
			want:    &mongo.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ctx, cancel := context.WithTimeout(tt.args.ctx, time.Second*30)
		defer cancel()
		t.Run(tt.name, func(t *testing.T) {
			client, err := dbutils.Conn(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer client.Disconnect(ctx)
		})
	}
}

func TestNewActiveIncidentServer(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *trainViewServer
		wantErr bool
	}{
		{
			name: "NewActiveIncidentServer",
			args: args{
				ctx: context.TODO(),
			},
			want:    &trainViewServer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTrainViewServer(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewActiveIncidentServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_ = got
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
			defer cancel()
			if result, err := got.EntriesMinutesAgo(ctx, 10); err == nil {
				fmt.Println(result)
			} else {
				t.Errorf("failure")
			}
		})
	}
}
