package apex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func CheckTempAndPh() error {
	status, err := GetStats()
	if err != nil {
		return err
	}
	fmt.Printf("Temp: %.1f  Ph: %.2f\n", status.Istat.Inputs[0].Value, status.Istat.Inputs[1].Value)
	return nil
}

func CheckTrident() error {
	status, err := GetStats()
	if err != nil {
		return err
	}
	alk, _ := findInput(status.Istat.Inputs, "alk")
	calc, _ := findInput(status.Istat.Inputs, "ca")
	mg, _ := findInput(status.Istat.Inputs, "mg")
	fmt.Printf("Alk: %.2f Calc %.0f Mg %.0f", alk.Value, calc.Value, mg.Value)
	return nil
}

func GetStats() (Status, error) {
	var status Status
	resp, err := http.Get("http://192.168.0.196/cgi-bin/status.json")
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
