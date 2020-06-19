package hardware

import (
	"fmt"
	"github.com/inktomi/squirrel/telemetry"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func setupPin() (rpio.Pin, error) {
	err := rpio.Open()
	if err != nil {
		fmt.Println("IO Error:", err)
		return rpio.Pin(BUZZER), err
	}

	var pin = rpio.Pin(BUZZER)
	pin.Output()

	return pin, nil
}

func closePin() {
	if err := rpio.Close(); err != nil {
		telemetry.ReportError(err, "Failed to close pin")
	}
}

func SingleBeep() error {
	pin, err := setupPin()
	if err != nil {
		return err
	}

	pin.Low()
	time.Sleep(500 * time.Millisecond)
	pin.High()

	closePin()

	return nil
}

func Alarm() error {
	pin, err := setupPin()
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		pin.Low() // Buzz
		time.Sleep(500 * time.Millisecond)
		pin.High() // Off
	}

	closePin()

	return nil
}
