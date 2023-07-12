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
	fmt.Printf("Name is %v and age is %v\n", stu.Name, stu.Age)
	stu.GetStatus()
	stu.UpdateName()
}

func (s Student) GetStatus() {
	fmt.Println("Is student active : ", s.Status)
}

func (u Student) UpdateName() {
	u.Name = "Bill"
	fmt.Println("The copy name is ", u.Name)
}
