package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Working with JSON Data")

	encodingJsonData(courses{"ReactJS", 234, "Android", "abc123", []string{"webdev", "js"}})
	// decodingJsonData("http://app-stage-01.wlink.com.np:30004/customers/test7777")
}

type courses struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"course_price"`
	Platform string   `json:"course_platform"`
	Password string   `json:"-"`
	Tag      []string `json:"tag,omitempty"`
}

func encodingJsonData(course courses) {
	jsonData, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	courseTest := []courses{
		{"Laravel", 344, "Development", "erw444", []string{"php", "laravel"}},
		{"Go", 455, "Development", "ert555", []string{"like-C", "compiler"}},
		{"Angular", 255, "Development", "rty44d", nil},
	}

	// courseJson, _ := json.MarshalIndent(courseTest, "", "\t")
	courseJson, _ := json.Marshal(courseTest)
	fmt.Println("This is coruse Json Marshal output")
	fmt.Printf("%s\n", courseJson)

	fmt.Println("Course Json Marshal output done")

	// fmt.Printf("%T", courseJson)

	decodingJsonData(courseJson)
}

func decodingJsonData(courseJsonData []byte) {
	response, err := http.Get("http://app-stage-01.wlink.com.np:30004/customers/test7777")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Println(response.Body)
	jsonDataFromWeb, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(body)
	// fmt.Printf("%T", body)
	var jsonData []courses
	checkValid := json.Valid(courseJsonData)
	// fmt.Println(checkValid)
	if checkValid {
		error := json.Unmarshal(courseJsonData, &jsonData)
		if error != nil {
			panic(error)
		}
		fmt.Printf("%#v\n", jsonData)
	}

	// some cases where we just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Println(myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("%v => %v \n ", key, value)
	}
}
