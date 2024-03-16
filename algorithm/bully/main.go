package main

import (
	"context"
	"time"

	"examples/algorithm/bully/pkg/mqtt"
)

type Node struct {
	ID       int
	IsOnline bool
}


type ElectionMessage struct {
	FromID int
}


func startElection() {

}

func findHigherNodes() {

}

func reveiveElectionMessage() {

}

func main() {
    config := mqtt.Config{
        Broker:   "mqtt.eclipseprojects.io:1883",
        ClientID: "mqtt-client-1",
        Topic:    "election",
    }

	// Initialize MessageHandler
	messageHandler := &mqtt.MessageHandler{}

	// Define a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Cancel the context to release resources when done

	// Start listening for messages with the defined context
	messageHandler.ConnectAndListen(ctx, config)
}
