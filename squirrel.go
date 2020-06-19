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
		defer func(client *telemetry.Adafruit) {
			if err := adafruitClient.Disconnect(); err != nil {
				telemetry.ReportError(err, "Failed to shut down Adafruit Client.")
			}
		}(adafruitClient)

		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			if err := adafruitClient.SendDataPoint(i); err != nil {
				telemetry.ReportError(err, "Failed to send telemetry data to Adafruit")
				break
			}
		}
	}

	//hardware.Buzz()
	//hardware.MonitorWeight(adafruitClient)
}
