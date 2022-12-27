package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJSON()
	decodeJSON()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "Learncodeonline.in", "xyx123", nil},
	}

	// package this data as JSON data

	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func decodeJSON() {
	jsonDataFromWeb := []byte(`
  {
	"coursename": "ReactJS Bootcamp",
	"Price": 299,
	"website": "Learncodeonline.in",
	"tags": ["web-dev","js"]
  }
  `)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!!")
	}

	// some cases where you just want to add data in key-value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
