package telemetry

import (
	"github.com/reef-pi/adafruitio"
	"os"
)

type Adafruit struct {
	adafruitClient *adafruitio.Client
	feed           string
	key            string
}

func CreateClient() *Adafruit {
	var key = os.Getenv("ADAFRUIT_IO_KEY")

	return &Adafruit{
		adafruitClient: adafruitio.NewClient(key),
		feed:           "load_cell_value",
		key:            key,
	}
}

func (af Adafruit) SendDataPoint(data int) error {
	telemetryData := adafruitio.Data{Value: data}

	err := af.adafruitClient.SubmitData(af.key, af.feed, telemetryData)
	return err
}
