package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ddasdsa443"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme is ", result.Scheme)
	fmt.Println("Host is ", result.Host)
	fmt.Println("Path is ", result.Path)
	fmt.Println("Raw query is ", result.RawQuery)
	fmt.Println("Port is ", result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=subham",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
