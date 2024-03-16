package mqtt

import (
    "context"
    "fmt"
    "log"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
    Broker   string
    ClientID string
    Topic    string
}

type MessageHandler struct{}

func (mh *MessageHandler) ConnectAndListen(ctx context.Context, config Config) {
    opts := MQTT.NewClientOptions().AddBroker(config.Broker).SetClientID(config.ClientID)
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    if token := client.Subscribe(config.Topic, 0, mh.messageHandler); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    fmt.Printf("Connected to MQTT broker. Listening to topic: %s\n", config.Topic)

    select {}
}

func (mh *MessageHandler) messageHandler(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Received message: %+v\n", msg)
}

func (mh *MessageHandler) PublishMessage(client MQTT.Client, topic string, message string) {
    token := client.Publish(topic, 0, false, message)
    token.Wait()
    if token.Error() != nil {
        log.Println("Error publishing message:", token.Error())
    } else {
        fmt.Printf("Published message '%s' to topic '%s'\n", message, topic)
    }
}
