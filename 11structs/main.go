package main

import "fmt"

// no inheritance in golang

type Student struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang")
	stu := Student{"Elon", "elonmusk@tesla.com", true, 30}
	fmt.Printf("%+v\n", stu)
	fmt.Printf("Name is %v and age is %v", stu.Name, stu.Age)
}
