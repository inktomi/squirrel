package main

import (
	"errors"
	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/inktomi/squirrel/hardware"
	"github.com/inktomi/squirrel/telemetry"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("squirrel.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	if err := hardware.Setup(); err != nil {
		log.Fatal("Failed to setup HX711", err)
	}

	if adafruitClient, err := telemetry.CreateClient(); err != nil {
		log.Fatal("Failed to setup & connect to MQTT Topic", err)
	} else {
		// Clean up MQTT
		defer func(client *telemetry.Adafruit) {
			if err := adafruitClient.Disconnect(); err != nil {
				log.Error(err, "Failed to shut down Adafruit Client.")
			}
		}(adafruitClient)

		// Clean up HX711
		defer func() {
			if err := hardware.Shutdown(); err != nil {
				log.Error(err, "Failed to shut down HX711")
			}
		}()

		// Set up the loop to track weights in
		var lastReported int64 = 0
		var movingAverage = movingaverage.New(1200)
		for {
			time.Sleep(100 * time.Millisecond)

			if weight, err := hardware.GetWeight(); err != nil {
				log.Error(err, "Failed to retrieve weight value")
			} else {
				// 10 weights per second
				// 600 weights per minute
				// 1200 weights for calibration
				movingAverage.Add(float64(weight))
				if movingAverage.Count() >= 100 {
					var zeroValue = movingAverage.Avg()

					var variance = math.Abs(zeroValue - float64(weight))
					//if variance > 500 {
					//	hardware.Alarm()
					//}

					if err := ReportWeightIfNeeded(&lastReported, adafruitClient, variance); err != nil {
						log.Error(err, "Failed to send telemetry data to Adafruit")
					}
				} else {
					// We're calibrating.
					log.WithFields(log.Fields{
						"weight":            weight,
						"calibration_count": movingAverage.Count(),
						"calibration_value": movingAverage.Avg(),
					}).Info("Added weight to calibration")
					//hardware.SingleBeep()
				}
			}
		}
	}
}

func ReportWeightIfNeeded(lastReported *int64, adafruitClient *telemetry.Adafruit, weight float64) error {
	var now = time.Now().Unix()
	var interval = now - *lastReported

	if interval > 20 {
		if err := adafruitClient.SendDataPoint(weight); err != nil {
			return err
		} else {
			log.WithFields(log.Fields{
				"weight":       weight,
				"now":          now,
				"lastReported": *lastReported,
				"interval":     interval,
			}).Info("Reported weight to adafruit.")
			lastReported = &now
		}
	} else {
		log.WithFields(log.Fields{
			"now":          now,
			"lastReported": *lastReported,
			"interval":     interval,
		}).Error("Reporting too fast.")

		return errors.New("tried to report too fast")
	}

	return nil
}
