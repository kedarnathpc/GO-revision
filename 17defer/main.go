package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("Hi")
	defer fmt.Println("Hey")
	fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
