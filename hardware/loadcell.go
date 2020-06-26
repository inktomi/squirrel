package hardware

import (
	"fmt"
	"github.com/MichaelS11/go-hx711"
	log "github.com/sirupsen/logrus"
)

var hx711chip *hx711.Hx711 = nil

func Setup() error {
	if err := hx711.HostInit(); err != nil {
		return err
	}

	if hx711chip != nil {
		if err := hx711chip.Shutdown(); err != nil {
			return err
		}
	}

	freshHx711, err := hx711.NewHx711(Hx711Clk, Hx711Data)
	if err != nil {
		return err
	} else {
		hx711chip = freshHx711
	}

	return nil
}

func Shutdown() error {
	if hx711chip == nil {
		return fmt.Errorf("cannot shutdown HX711, no chip configured")
	}

	if err := hx711chip.Shutdown(); err != nil {
		log.Error(err, "Failed to shut down HX711")
		return err
	}

	return nil
}

func GetWeight() (int, error) {
	if err := hx711chip.Reset(); err != nil {
		log.Error(err, "Reset HX711 had an error")
		return 0, err
	}

	var data int
	if rawData, err := hx711chip.ReadDataRaw(); err != nil {
		log.Error(err, "ReadDataRaw had an error")
		return 0, err
	} else {
		data = rawData
	}

	fmt.Println(data)

	return data, nil
}
