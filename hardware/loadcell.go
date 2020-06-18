package hardware

import (
	"fmt"
	"github.com/MichaelS11/go-hx711"
	"github.com/inktomi/squirrel/telemetry"
	"time"
)

func MonitorWeight(adafruit *telemetry.Adafruit) {
	hostInitErr := hx711.HostInit()
	if hostInitErr != nil {
		fmt.Println("HostInit error:", hostInitErr)
		return
	}

	hx711chip, err := hx711.NewHx711(Hx711Clk, Hx711Data)
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

		err := adafruit.SendDataPoint(data)
		if err != nil {
			fmt.Println("Error sending data to adafruit:", err)
		}

		fmt.Println(data)
	}
}
