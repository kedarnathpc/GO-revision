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
	fmt.Println("Welcome to JSON !!!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "LearnCodeOnline.in", "def123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "ghi123", nil},
	}

	// package this data as json data

	finaljson, err := json.MarshalIndent(lcoCourses, "", "\t")
	Checkerror(err)
	fmt.Printf("%s\n", finaljson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "React Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var onlineData map[string]interface{}
	json.Unmarshal(jsonData, &onlineData)
	fmt.Printf("%#v\n\n\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("%v : %v and type: %T\n", k, v, v)
	}
}

func Checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
