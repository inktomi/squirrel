package main

import (
	"github.com/reef-pi/adafruitio"
	"os"
)

func configureTelemetry() *adafruitio.Client {
	adafruitClient := adafruitio.NewClient(os.Getenv("ADAFRUIT_IO_KEY"))

	return adafruitClient
}