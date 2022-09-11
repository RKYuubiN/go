package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	Count    int    `json:"count"`
	Ackmode  string `json:"ackmode"`
	Encoding string `json:"encoding"`
	Truncate int    `json:"truncate"`
}

type Properties struct {
	DeliveryMode int    `json:"delivery_mode"`
	ContentType  string `json:"content_type"`
}

type Response struct {
	PayloadBytes    int        `json:"payload_bytes"`
	Redelivered     bool       `json:"redelivered"`
	Exchange        string     `json:"exchange"`
	RoutingKey      string     `json:"routing_key"`
	MessageCount    int        `json:"message_count"`
	Properties      Properties `json:"properties"`
	Payload         string     `json:"payload"`
	PayloadEncoding string     `json:"payload_encoding"`
}

func main() {

	data := Payload{100, "ack_requeue_true", "auto", 50000}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err", err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:15672/api/queues/%2f/something-queue/get", body)
	if err != nil {
		fmt.Println("err", err)
	}

	req.SetBasicAuth("guest", "guest")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("Using map[string]interface{} for response")
	fmt.Println()

	var queueMessage []map[string]interface{}

	err = json.Unmarshal(bodyResp, &queueMessage)

	if err != nil {
		fmt.Println(err)
	}

	for _, value := range queueMessage {
		fmt.Printf("%v \n", value["payload"])
	}

	fmt.Println()
	fmt.Println("Using structure for response")
	var queueMessageStruct []Response

	err = json.Unmarshal(bodyResp, &queueMessageStruct)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range queueMessageStruct {
		fmt.Println(value.Payload)
	}
	defer resp.Body.Close()

}
