package apex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/troycornwall/apexWatch/Config"
	"github.com/troycornwall/apexWatch/slack"
	"io"
	"net/http"
)

func CheckTempAndPh(cfg Config.Config) error {
	status, err := GetStats(cfg)
	if err != nil {
		return err
	}
	temp, _ := findInput(status.Istat.Inputs, "Temp")
	ph, _ := findInput(status.Istat.Inputs, "pH")

	log := fmt.Sprintf("Temp: %.1f  Ph: %.2f", temp.Value, ph.Value)
	fmt.Println(log)

	if temp.Value >= cfg.Limits.Temp.High || temp.Value <= cfg.Limits.Temp.Low {
		s := fmt.Sprintf("Temp out of range! Temp: %.1f", temp.Value)
		fmt.Println(s)
		slack.SendLogs(cfg.SlackWebhook, s)
	}
	if ph.Value >= cfg.Limits.Ph.High || ph.Value <= cfg.Limits.Ph.Low {
		s := fmt.Sprintf("PH out of range! Ph: %.2f", ph.Value)
		fmt.Println(s)
		slack.SendLogs(cfg.SlackWebhook, s)
	}
	return nil
}

func CheckTrident(cfg Config.Config) error {
	status, err := GetStats(cfg)
	if err != nil {
		return err
	}
	alk, _ := findInput(status.Istat.Inputs, "alk")
	calc, _ := findInput(status.Istat.Inputs, "ca")
	mg, _ := findInput(status.Istat.Inputs, "mg")
	s := fmt.Sprintf("Alk: %.2f Calc %.0f Mg %.0f", alk.Value, calc.Value, mg.Value)
	fmt.Println(s)
	slack.SendLogs(cfg.SlackWebhook, s)
	return nil
}

func GetStats(cfg Config.Config) (Status, error) {
	var status Status
	url := fmt.Sprintf("http://%s/cgi-bin/status.json", cfg.Hostname)
	resp, err := http.Get(url)
	if err != nil {
		return status, err
	}
	defer resp.Body.Close()

	jsonBytes, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return status, err
	}

	json.Unmarshal(jsonBytes, &status)
	return status, nil
}

func findInput(inputs []Input, key string) (Input, error) {
	for i := 0; i < len(inputs); i++ {
		if inputs[i].Type == key {
			return inputs[i], nil
		}
	}
	fmt.Printf("couldn't find '%s'\n", key)
	return Input{}, errors.New("couldn't find input")
}
