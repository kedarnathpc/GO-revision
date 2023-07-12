package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request in golang !!!")

	response, err := http.Get(url)
	checkerror(err)

	fmt.Printf("The response if of type : %T\n", response)
	defer response.Body.Close() // we should close the request, coders responsibility

	databytes, err := ioutil.ReadAll(response.Body)
	checkerror(err)
	content := string(databytes)
	fmt.Println("The body of the response is : ", content)
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
