package main

import (
	"github.com/go-co-op/gocron"
	"github.com/troycornwall/apexWatch/apex"
	"time"
)

func main() {
	t, _ := time.LoadLocation("NZ")
	s := gocron.NewScheduler(t)

	s.Every(2).Minute().Do(func() {
		apex.CheckTempAndPh()
	})

	s.Cron("30 0,3,6,9,12,15,18,21 * * *").Do(func() {
		apex.CheckTrident()
	})

	s.StartBlocking()
}
