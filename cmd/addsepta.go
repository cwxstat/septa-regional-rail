package cmd

import (
	"context"
	"github.com/cwxstat/septa-regional-rail/constants"
	"github.com/cwxstat/septa-regional-rail/hydrate"
	"github.com/cwxstat/septa-regional-rail/metrics"
	"github.com/cwxstat/septa-regional-rail/trainview"
	"log"
	"time"
)

func AddTrainView() error {
	metrics.RootStartLoops()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ase, err := trainview.NewTrainViewServer(ctx)
	if err != nil {
		log.Println(err)
		time.Sleep(constants.ErrorBackoff)
		return err
	}
	ase.DatabaseCollection("Septa", "trainView")

	page, err := hydrate.Grab(constants.TRAINVIEW)
	if err != nil {
		log.Println(err)
		time.Sleep(constants.ErrorBackoff)
		return err
	}
	trains, err := hydrate.Hydrate(page)
	if err != nil {
		log.Println(err)
		time.Sleep(constants.ErrorBackoff)
		return err
	}

	data := &trainview.ActiveSeptaEntry{
		MainWebPage: string(page),
		TrainView:   *trains,
		Message:     "",
		TimeStamp:   time.Now(),
	}

	err = ase.AddEntry(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ase.AddEntry(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := ase.Disconnect(ctx); err != nil {
		log.Println("as.Disconnect: ", err)
	}
	log.Println("entry added")
	metrics.RootProcessedLoops()
	return nil
}
