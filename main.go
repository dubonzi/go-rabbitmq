package main

import (
	"fmt"
	"os"

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

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(q)

	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello from Client"),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
