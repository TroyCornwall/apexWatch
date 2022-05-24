package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/troycornwall/apexWatch/Config"
	"github.com/troycornwall/apexWatch/apex"
	"os"
	"time"
)

func main() {
	var cfg Config.Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println("Failed to load config")
		os.Exit(2)
	}

	t, _ := time.LoadLocation(cfg.Timezone)
	s := gocron.NewScheduler(t)

	s.Every(2).Minute().Do(func() {
		err := apex.CheckTempAndPh(cfg)
		handleError(err)
	})

	s.Cron("30 0,3,6,9,12,15,18,21 * * *").Do(func() {
		err := apex.CheckTrident(cfg)
		handleError(err)
	})

	s.StartBlocking()
}

func handleError(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}
