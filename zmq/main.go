package main

import (
	"fmt"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	// Create a SUB socket
	subscriber, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		panic(err)
	}
	defer subscriber.Close()

	// Connect to the Vert.x publisher
	err = subscriber.Connect("tcp://localhost:6000")
	if err != nil {
		panic(err)
	}

	// Subscribe to all messages
	err = subscriber.SetSubscribe("")
	if err != nil {
		panic(err)
	}

	fmt.Println("Subscriber connected and waiting for messages...")

	for {
		// Receive a message
		msg, err := subscriber.Recv(0)
		if err != nil {
			fmt.Println("Error receiving message:", err)
		} else {
			fmt.Println("Received message:", msg)
		}
		// Simulate processing delay
		time.Sleep(500 * time.Millisecond)
	}
}
