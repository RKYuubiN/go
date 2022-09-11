package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers
var signals = []string{"signals"}

var mut sync.Mutex //should be pointers

func main() {
	fmt.Println("Working with Go routines")
	// go greeter("Shraddha")
	// greeter("Ayusha")

	websiteList := []string{
		"http://app-stage-01.wlink.com.np:30004/customers/test7777",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCodes(web)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(name string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println("Hello ", name)
// 	}
// }

func getStatusCodes(endPoint string) {
	defer wg.Done()
	response, _ := http.Get(endPoint)
	mut.Lock()
	signals = append(signals, endPoint)
	mut.Unlock()
	fmt.Printf("Response code for %v is %v\n", endPoint, response.StatusCode)
}
