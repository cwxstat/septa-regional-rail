package hydrate

import (
	"encoding/json"
	"github.com/cwxstat/septa-regional-rail/get"
)

type TrainView []struct {
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Trainno     string  `json:"trainno"`
	Service     string  `json:"service"`
	Dest        string  `json:"dest"`
	Currentstop string  `json:"currentstop"`
	Nextstop    string  `json:"nextstop"`
	Line        string  `json:"line"`
	Consist     string  `json:"consist"`
	Heading     float64 `json:"heading"`
	Late        int     `json:"late"`
	Source      string  `json:"SOURCE"`
	Track       string  `json:"TRACK"`
	TrackChange string  `json:"TRACK_CHANGE"`
}

func Hydrate(b []byte) (*TrainView, error) {

	trainView := &TrainView{}
	if err := json.Unmarshal(b, &trainView); err != nil {
		return nil, err
	}
	return trainView, nil
}

func Grab(url string) ([]byte, error) {
	s, err := get.Get(url)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}
