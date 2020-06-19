package telemetry

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

const feedName = "load-cell-value"

type Adafruit struct {
	client mqtt.Client
	topic  string
	errors string
	feed   string
	key    string
}

var handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("MQTT TOPIC: %s, MSG: %s\n", msg.Topic(),  msg.Payload())
}

func CreateClient() (*Adafruit, error) {
	var key = os.Getenv("ADAFRUIT_IO_KEY")
	var username = os.Getenv("ADAFRUIT_IO_USER")

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

	if token := client.Subscribe(username+"/feeds/"+feedName, 0, handler); token.Wait() && token.Error() != nil {
		ReportError(token.Error(), "Failed to subscribe to MQTT Feed")
		return &Adafruit{}, token.Error()
	}

	if token := client.Subscribe(username+"/errors", 0, handler); token.Wait() && token.Error() != nil {
		ReportError(token.Error(), "Failed to subscribe to MQTT Error Feed")
		return &Adafruit{}, token.Error()
	}

	return &Adafruit{
			client: client,
			topic: username+"/feeds/"+feedName,
			errors: username+"/errors",
			feed: feedName,
			key:  key,
		}, nil
}

func (af Adafruit) Disconnect() error {
	if token := af.client.Unsubscribe(af.feed, af.errors); token.Wait() && token.Error() != nil {
		ReportError(token.Error(), "Failed to unsubscribe.")
		return token.Error()
	}

	af.client.Disconnect(250)

	return nil
}

func (af Adafruit) SendDataPoint(data int) error {
	var payload = fmt.Sprintf("%v", data)

	if token := af.client.Publish(af.topic, 0, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
