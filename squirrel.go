package main

import (
	"github.com/inktomi/squirrel/hardware"
	"github.com/inktomi/squirrel/telemetry"
)

func main() {
	telemetry.ConfigureBugsnag()
	adafruitClient := telemetry.CreateClient()

	hardware.Buzz()
	hardware.MonitorWeight(adafruitClient)
}
