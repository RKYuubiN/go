package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to rabbit mq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"some-queue",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	// body := "hello world"

	body := bodyFrom(os.Args)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf("[x] Sent %s \n", body)

}

func bodyFrom(args []string) string {
	// rand.Seed(time.Now().Unix())
	// s := []string{"spaghetti", "broccoli", "pasta", "cheese", "dark chocolate"}

	var dummyData string
	if (len(args) < 2) || os.Args[1] == "" {
		// dummyData = s[rand.Int(len(s))]
		dummyData = "cheese"
	} else {
		dummyData = strings.Join(args[1:], " ")
	}
	return dummyData
}
