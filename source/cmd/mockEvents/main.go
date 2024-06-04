package mockevents

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"source/cmd/eventSource"
	"source/configs"
	"source/models"
	"time"
)

func GenerateEvents(ctx context.Context, eventChan chan<- models.UserInteractionData) {
	ticker := time.NewTicker(configs.Event_IntialEventTime * 
                                    configs.Event_InitalEventTimeUnit)
	var resetTicker = func() {
		t := rand.Intn(configs.Event_NextEventTimeRange)
		// if t <= 0 it produce panic when reset
		if t <= 0 {
			t = 5
		}

		ticker.Reset(time.Duration(t) * configs.Event_NextEventTimeUnit)
		fmt.Printf("\n------Next event in %d second-----\n", t)
	}

outer:
	for {
		select {
		case <-ticker.C:
			data, err := eventSource.NewUserInteractionData()
			if err != nil {
				slog.Error("Unable to souce the event", "Details", err.Error())
				resetTicker()
				continue
			}
			eventChan <- *data
			resetTicker()
		case <-ctx.Done():
			break outer
		}
	}
	ticker.Stop()
	close(eventChan)
	slog.Info("Generating Events is Done")

}
