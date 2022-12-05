package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename":"Let's go with golang",
		"price": 0,
		"platform":"learncodeonline.in"
	}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func PerformFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "subham")
	data.Add("lastname", "bharadwaz")
	data.Add("email", "subhm@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
