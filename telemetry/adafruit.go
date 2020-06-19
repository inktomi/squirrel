package telemetry

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

const feedName = "load_cell_value"

type Adafruit struct {
	client mqtt.Client
	topic  string
	feed   string
	key    string
}

var handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func CreateClient() (*Adafruit, error) {
	var key = os.Getenv("ADAFRUIT_IO_KEY")
	var username = "inktomi" // TODO: ENV value.

	options := mqtt.NewClientOptions().
		AddBroker("wss://io.adafruit.com:443").
		SetUsername(username).
		SetPassword(key).
		SetKeepAlive(2 * time.Second).
		SetDefaultPublishHandler(handler).
		SetPingTimeout(1 * time.Second).
		SetAutoReconnect(true)

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		ReportError(token.Error(), "Failed to connect to Adafruit MQTT broker")
		return &Adafruit{}, token.Error()
	}

	if token := client.Subscribe(username+"/feeds/"+feedName, 0, nil); token.Wait() && token.Error() != nil {
		ReportError(token.Error(), "Failed to subscribe to MQTT Feed")
		return &Adafruit{}, token.Error()
	}

	return &Adafruit{
			client: client,
			topic: username+"/feeds/"+feedName,
			feed: feedName,
			key:  key,
		}, nil
}

func (af Adafruit) GetOrCreateFeed() error {

	//var _, feedError = client.GetFeed(user.Name, feedName)
	//if feedError != nil {
	//	ReportError(feedError, "Could not get feed with name: " + feedName)
	//
	//	var createErr = client.CreateFeed(user.Name, adafruitio.Feed{Name: feedName})
	//	if createErr != nil {
	//		ReportError(createErr, "Could not create feed.")
	//	}
	//}

	return nil
}

func (af Adafruit) SendDataPoint(data int) error {
	if token := af.client.Publish(af.topic, 0, false, data); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
