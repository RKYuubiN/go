
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Handling URLs and working with http in go")
	const myUrl string = "http://app-stage-01.wlink.com.np:30004/customers/test7777"
	const myPostURL string = "https://simple-books-api.glitch.me/api-clients"
	const myPostURLFormData string = "https://jsonplaceholder.typicode.com/posts"
	// parsing
	result, err := url.Parse(myUrl)
	checkNilError(err)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawPath)
	fmt.Println(result.Port())

	partsOfURL := &url.URL{
		Scheme: "http",
		Host:   "app-stage-01.wlink.com.np:30004",
		Path:   "/customers/test7777",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)

	// GET Request
	// makeGetRequest(myUrl)

	// POST Request
	// makePostRequest(myPostURL)
	makePostRequestForm(myPostURLFormData)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// GET Request
func makeGetRequest(requestUrl string) {
	resp, err := http.Get(requestUrl)
	checkNilError(err)
	defer resp.Body.Close()
	var responseString strings.Builder
	body, err := io.ReadAll(resp.Body)
	responseString.Write(body)
	checkNilError(err)
	fmt.Println("Status code: ", resp.Status)
	fmt.Println("Content Length", resp.ContentLength)
	fmt.Println(responseString.String())
}

// POST Request
func makePostRequest(requestUrl string) {
	payload := strings.NewReader(`
	{
			"clientName":"sarah123",
			"clientEmail":"sarah123@example.dev"
		}
		`)

	fmt.Println(payload)
	// byteData, err := json.Marshal(payload)
	// checkNilError(err)
	// fmt.Println(byteData)

	resp, err := http.Post(requestUrl, "application/json", payload)
	checkNilError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var responseString strings.Builder
	responseString.Write(body)
	fmt.Println(responseString.String())
}

// POST Request with form-data
func makePostRequestForm(requestUrl string) {
	values := url.Values{}
	values.Set("title", "foo")
	values.Set("body", "bar")
	values.Set("userId", "2")
	response, err := http.PostForm(requestUrl, values)
	checkNilError(err)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	var responseString strings.Builder
	responseString.Write(body)
	fmt.Println(responseString.String())

}
