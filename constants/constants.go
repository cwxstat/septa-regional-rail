package constants

import "time"

var (
	TRAINVIEW    string = "https://www3.septa.org/hackathon/TrainView"
	ErrorBackoff        = 3 * time.Second
	RefreshRate         = time.Second * 10
)
