package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	//"github.com/MichaelS11/go-hx711"
)

import "github.com/bugsnag/bugsnag-go"

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: os.Getenv("BUGSNAG_KEY"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/inktomi/squirrel"},
	})

	err := rpio.Open()
	if err != nil {
		fmt.Println("IO Error:", err)
		return
	}

	pin := rpio.Pin(6)

	pin.Output()

	for i := 0; i < 1000; i++ {
		time.Sleep(500 * time.Microsecond)
		pin.High()
		time.Sleep(200 * time.Microsecond)
		pin.Low()
	}

	//err := hx711.HostInit()
	//if err != nil {
	//	fmt.Println("HostInit error:", err)
	//	return
	//}
	//
	//hx711chip, err := hx711.NewHx711("GPIO5", "GPIO4")
	//if err != nil {
	//	fmt.Println("NewHx711 error:", err)
	//	return
	//}
	//
	//defer hx711chip.Shutdown()
	//
	//err = hx711chip.Reset()
	//if err != nil {
	//	fmt.Println("Reset error:", err)
	//	return
	//}
	//
	//var data int
	//for i := 0; i < 10000; i++ {
	//	time.Sleep(200 * time.Microsecond)
	//
	//	data, err = hx711chip.ReadDataRaw()
	//	if err != nil {
	//		fmt.Println("ReadDataRaw error:", err)
	//		continue
	//	}
	//
	//	fmt.Println(data)
	//}

}
