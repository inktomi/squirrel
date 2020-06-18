package hardware

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func Buzz() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("IO Error:", err)
		return
	}

	pin := rpio.Pin(BUZZER)

	pin.Output()

	fmt.Println("Starting to buzz")
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		pin.Low() // Buzz
		time.Sleep(500 * time.Millisecond)
		pin.High() // Off
	}

	closeError := rpio.Close()
	if closeError != nil {
		fmt.Println("Failed to close:", closeError)
	}
}
