package scheduler

import (
	"time"

	gocron "github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler = nil

func GetScheduler() *gocron.Scheduler {

	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.Local)
	}
	return scheduler

}
