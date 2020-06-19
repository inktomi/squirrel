package telemetry

import (
	"fmt"
	"os"
)
import "github.com/bugsnag/bugsnag-go"

func ConfigureBugsnag() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: os.Getenv("BUGSNAG_KEY"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/inktomi/squirrel"},
	})
}

func ReportError(err error, message string) {
	bugsnagErr := bugsnag.Notify(err, message)

	if bugsnagErr != nil {
		fmt.Printf("Failed to report error: `%v` because of `%v`", err, bugsnagErr)
	}
}
