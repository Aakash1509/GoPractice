package main

import (
	"fmt"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	publisher, err := zmq4.NewSocket(zmq4.PUB)
	if err != nil {
		panic(err)
	}
	defer publisher.Close()

	publisher.Bind("tcp://*:6000")
	fmt.Println("Publisher started...")

	for count := 0; ; count++ {
		msg := fmt.Sprintf("Message %d", count)
		fmt.Println("Publishing:", msg)
		publisher.Send(msg, 0)
		time.Sleep(1 * time.Second)
	}
}
