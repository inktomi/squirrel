package main

import (
	"github.com/inktomi/squirrel/telemetry"
	"log"
	"time"
)

func main() {
	telemetry.ConfigureBugsnag()
	if adafruitClient, err := telemetry.CreateClient(); err != nil {
		log.Panicf("Failed to setup & connect to MQTT Topic: %v", err)
	} else {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			err := adafruitClient.SendDataPoint(i)
			if err != nil {
				telemetry.ReportError(err, "Failed to send telemetry data to Adafruit")
			}
		}
	}

	//hardware.Buzz()
	//hardware.MonitorWeight(adafruitClient)
}
