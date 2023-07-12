package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dfd432dfe"

func main() {
	fmt.Println("Welcome to handling URLs in golang !!!")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println("The scheme is : ", result.Scheme)
	fmt.Println("The host is : ", result.Host)
	fmt.Println("The path is : ", result.Path)
	fmt.Println("The port is : ", result.Port())
	fmt.Println("The rawquery is : ", result.RawQuery)

	// analysing the rawQuery
	qparams := result.Query()
	fmt.Printf("The type of query params is : %T\n", qparams)
	for key, val := range qparams {
		fmt.Printf("key : %v value : %v\n", key, val)
	}

	//combining the parts of URL

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsofURL.String()
	fmt.Println("Another URL is : ", anotherURL)
}
