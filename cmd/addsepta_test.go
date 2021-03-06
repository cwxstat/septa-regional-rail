package cmd

import "testing"

func TestAddTrainView(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Look for error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddTrainView(); (err != nil) != tt.wantErr {
				t.Errorf("AddTrainView() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
