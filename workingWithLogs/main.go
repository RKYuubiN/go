package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	fmt.Println("working with logs")

	// var logger log.Logger
	logger := log.NewJSONLogger(os.Stdout)

	text := map[string]string{
		`msg`:  `hello`,
		`msg2`: `hello2`,
		`msg3`: `hello3`,
	}

	jsonText, _ := json.Marshal(text)

	var decodedJson map[string]string
	err := json.Unmarshal(jsonText, &decodedJson)
	if err != nil {
		fmt.Println(err)
	}

	level.Info(logger).Log("message 1", "here here", "msg2", "hello", "payload", decodedJson)

}
