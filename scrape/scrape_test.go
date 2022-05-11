package scrape

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		url    string
		client []*http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Quick Smoke Test",
			args: args{
				url: "https://www3.septa.org/hackathon/TrainView",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 0 {
				fmt.Printf("output:\n%v\n", got)
			}
		})
	}
}
