package main

import (
	"fmt"
	"github.com/MichaelS11/go-hx711"
	"github.com/inktomi/squirrel/hardware"
)

// O hook weight: 104.6 grams  (use as nothing hanging weight)
// Hummingbird Feeder (Full): 1091 grams

/**
  AdjustZero should be set to: 16383
  AdjustScale should be set to a value between 0.000000 and 41.110739
*/

func main() {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711chip, err := hx711.NewHx711(hardware.Hx711Clk, hardware.Hx711Data)
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}

	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711chip.SetGain(128)

	var weight1 float64
	var weight2 float64

	weight1 = 104.6
	weight2 = 1195.6

	hx711chip.GetAdjustValues(weight1, weight2)
}
