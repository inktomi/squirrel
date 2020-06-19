package main

import (
	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/inktomi/squirrel/hardware"
	"github.com/inktomi/squirrel/telemetry"
	"log"
	"math"
	"time"
)

func main() {
	telemetry.ConfigureBugsnag()
	if err := hardware.Setup(); err != nil {
		log.Panicf("Failed to setup HX711: %v", err)
	}

	if adafruitClient, err := telemetry.CreateClient(); err != nil {
		log.Panicf("Failed to setup & connect to MQTT Topic: %v", err)
	} else {
		// Clean up MQTT
		defer func(client *telemetry.Adafruit) {
			if err := adafruitClient.Disconnect(); err != nil {
				telemetry.ReportError(err, "Failed to shut down Adafruit Client.")
			}
		}(adafruitClient)

		// Clean up HX711
		defer func() {
			if err := hardware.Shutdown(); err != nil {
				telemetry.ReportError(err, "Failed to shut down HX711")
			}
		}()

		// Set up the loop to track weights in
		var lastReported = time.Now()
		var movingAverage = movingaverage.New(1200)
		for {
			time.Sleep(100 * time.Millisecond)

			if weight, err := hardware.GetWeight(); err != nil {
				telemetry.ReportError(err, "Failed to retrieve weight value")
			} else {
				// 10 weights per second
				// 600 weights per minute
				// 1200 weights for calibration
				movingAverage.Add(float64(weight))
				if movingAverage.Count() >= 1200 {
					var zeroValue = movingAverage.Avg()

					var variance = math.Abs(zeroValue - float64(weight))
					//if variance > 500 {
					//	hardware.Alarm()
					//}

					if err := reportWeightIfNeeded(lastReported, adafruitClient, variance); err != nil {
						telemetry.ReportError(err, "Failed to send telemetry data to Adafruit")
					}
				} else {
					// We're calibrating.
					//hardware.SingleBeep()
				}
			}
		}
	}
}

func reportWeightIfNeeded(lastReported time.Time, adafruitClient *telemetry.Adafruit, weight float64) error {
	if time.Now().Sub(lastReported) >= 3*time.Second {
		if err := adafruitClient.SendDataPoint(weight); err != nil {
			return err
		} else {
			lastReported = time.Now()
		}
	}

	return nil
}
