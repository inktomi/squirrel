package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// import "github.com/stianeikeland/go-rpio"
import "github.com/bugsnag/bugsnag-go"

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: os.Getenv("BUGSNAG_KEY"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/inktomi/squirrel"},
	})

	err := bugsnag.Notify(fmt.Errorf("test error"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	log.Print("Finished")

}
