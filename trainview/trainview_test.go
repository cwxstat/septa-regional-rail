package trainview

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cwxstat/septa-regional-rail/constants"
	"github.com/cwxstat/septa-regional-rail/dbutils"
	"github.com/cwxstat/septa-regional-rail/hydrate"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestFull(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	as, err := NewTrainViewServer(ctx)
	if err != nil {
		t.FailNow()
	}

	as.DatabaseCollection("testSepta", "trainView")

	page, err := hydrate.Grab(constants.TRAINVIEW)
	if err != nil {
		t.FailNow()
	}
	trainview, err := hydrate.Hydrate(page)
	if err != nil {
		t.FailNow()
	}
	data := &ActiveSeptaEntry{
		MainWebPage: string(page),
		TrainView:   *trainview,
		Message:     "",
		TimeStamp:   time.Now(),
	}

	err = as.AddEntry(ctx, data)
	if err != nil {
		t.FailNow()
	}

}

func Grab(s string) {
	panic("unimplemented")
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

func TestNewTrainViewServer(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *trainViewServer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTrainViewServer(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTrainViewServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrainViewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
