package main

import (
	"fmt"
	"github.com/MichaelS11/go-hx711"
	"github.com/reef-pi/adafruitio"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func main() {
	configureBugsnag()
	client := configureTelemetry()

	err := rpio.Open()
	if err != nil {
		fmt.Println("IO Error:", err)
		return
	}

	pin := rpio.Pin(6)

	pin.Output()

	fmt.Println("Starting to buzz")
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		pin.Low() // Buzz
		time.Sleep(500 * time.Millisecond)
		pin.High() // Off
	}

	fmt.Println("Done. Good day.")

	closeError := rpio.Close()
	if closeError != nil {
		fmt.Println("Failed to close:", closeError)
	}

	hostInitErr := hx711.HostInit()
	if hostInitErr != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711chip, err := hx711.NewHx711("GPIO5", "GPIO4")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}

	defer hx711chip.Shutdown()

	err = hx711chip.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}

	var data int
	for i := 0; i < 10000; i++ {
		time.Sleep(2 * time.Second)

		data, err = hx711chip.ReadDataRaw()
		if err != nil {
			fmt.Println("ReadDataRaw error:", err)
			continue
		}

		telemetryData := adafruitio.Data { Value: data }
		feed := "load_cell_value"

		err := client.SubmitData(os.Getenv("ADAFRUIT_IO_KEY"), feed, telemetryData)
		if err != nil {
			fmt.Println("Submit telemetry error:", err)
		}

		fmt.Println(data)
	}

}
