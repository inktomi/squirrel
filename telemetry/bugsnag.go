package telemetry

import "os"
import "github.com/bugsnag/bugsnag-go"

func ConfigureBugsnag() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: os.Getenv("BUGSNAG_KEY"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/inktomi/squirrel"},
	})

}
