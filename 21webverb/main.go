package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb !!!")
	// PerformGetRequest()
	// PerformPostjsonRequest()
	PerformPostformRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	Checkerror(err)
	defer response.Body.Close()
	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostjsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	Checkerror(err)

	content, err := ioutil.ReadAll(response.Body)
	Checkerror(err)
	defer response.Body.Close()
	fmt.Println(string(content))
}

func PerformPostformRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	Checkerror(err)
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	Checkerror(err)
	fmt.Println(string(content))
}

func Checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
