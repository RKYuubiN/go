package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")

	var someValue = flag.String("fruit", "Apple", "Favorite Fruits")
	flag.Parse()
	fmt.Println(*nFlag)
	fmt.Println(*someValue)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	})

	http.HandleFunc("/testValue", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Something something testing")
	})

	http.ListenAndServe(":8090", nil)
}
