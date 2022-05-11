package hydrate

import (
	"testing"

	"github.com/cwxstat/septa-regional-rail/constants"
)

func TestGrab(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "One Pull",
			args: args{
				url: constants.TRAINVIEW,
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grab(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grab() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result, err := Hydrate(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hydrate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(*result) < 4 {
				t.Errorf("result not valid")
				return
			}

		})
	}
}
