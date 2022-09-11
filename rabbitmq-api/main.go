package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Response struct {
	PayloadBytes int    `json:"payload_bytes"`
	Redelivered  bool   `json:"redelivered"`
	Exchange     string `json:"exchange"`
	RoutingKey   string `json:"routing_key"`
	MessageCount int    `json:"message_count"`
	Properties   struct {
		DeliveryMode int    `json:"delivery_mode"`
		ContentType  string `json:"content_type"`
	} `json:"properties"`
	Payload         string `json:"payload"`
	PayloadEncoding string `json:"payload_encoding"`
}

type Payload struct {
	Count    int    `json:"count"`
	AckMode  string `json:"ackmode"`
	Encoding string `json:"encoding"`
	Truncate int    `json:"truncate"`
}

func main() {
	fmt.Println("Working with rabbitmq api")

	r := chi.NewRouter()

	r.Get("/rabbitmq/queue/messages", getMessages)

	// Get response from rabbitmq http api package
	data := Payload{1000, "ack_requeue_true", "auto", 50000}

	payloadBytes, err := json.Marshal(data)

	FailOnError(err)

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:15672/api/queues/%2f/some-queue/get", body)
	FailOnError(err)

	req.SetBasicAuth("guest", "guest")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	FailOnError(err)

	bodyResp, err := io.ReadAll(resp.Body)
	FailOnError(err)

	var queueResponsePayload []Response
	err = json.Unmarshal(bodyResp, &queueResponsePayload)
	FailOnError(err)

	fmt.Println(queueResponsePayload)

	// for data above count 1000
	if queueResponsePayload[0].MessageCount > 1000 {
		fmt.Println("start a new response")
	}
}

func getMessages(w http.ResponseWriter, r *http.Request) {

}

func FailOnError(err error) {
	if err != nil {
		log.Printf("Errors are %+v", err)
	}
}
