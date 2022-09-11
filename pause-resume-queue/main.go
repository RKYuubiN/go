package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func startQueue(ch *amqp.Channel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Has("queuename") && query.Has("consumerkey") {
			queue, _ := ch.QueueInspect(query.Get("queuename"))
			if queue.Consumers > 0 {
				json.NewEncoder(w).Encode("Error. Queue Consumption already Started")
			} else {
				_, err := ch.Consume(
					query.Get("queuename"),
					query.Get("consumerkey"),
					true,
					false,
					false,
					false,
					nil,
				)
				failOnErr("Failing on consume", err)
				json.NewEncoder(w).Encode("Status 200 OK. Queue Consumption Started")
			}
		} else {
			json.NewEncoder(w).Encode("Status 400 Bad Request. Please provide queuename and consumerkey in query string")
		}
	}

}

func stopQueue(ch *amqp.Channel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Status OK. Queue Consumption Stopped")
		query := r.URL.Query()
		if query.Has("consumerkey") {
			err := ch.Cancel(query.Get("consumerkey"), false)
			failOnErr("Error on canceling channel", err)
		} else {
			json.NewEncoder(w).Encode("Status 400 Bad Request. Please provide consumerkey in query string")
		}
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr("Error from opening connections", err)
	defer conn.Close()
	fmt.Println("Opened a connection")

	ch, err := conn.Channel()
	failOnErr("Error from opening a channel", err)

	fmt.Println("Opened a channel")
	defer ch.Close()

	http.HandleFunc("/start", startQueue(ch))
	http.HandleFunc("/stop", stopQueue(ch))

	fmt.Println("Listening at port 8080")

	http.ListenAndServe(":8080", nil)
}

func failOnErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s %d", message, err)
	}
}
