package hydrate

import (
	"encoding/json"
	"github.com/cwxstat/septa-regional-rail/get"
	"github.com/cwxstat/septa-regional-rail/trainview"
)

func Hydrate(b []byte) (*trainview.TrainView, error) {

	trainView := &trainview.TrainView{}
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
