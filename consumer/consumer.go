package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Connecting to RMQ Instance")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ch.Close()

	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, os.Kill)

	go func() {
		for msg := range msgs {
			fmt.Printf("New message: %s\n", msg.Body)
		}
	}()

	fmt.Println("{*} Listening for messages...")
	<-stop
}
