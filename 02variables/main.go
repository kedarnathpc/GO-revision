package main

import "fmt"

const Pi = 3.14 //public

func main() {

	// types

	fmt.Println("Hello there")
	var name string = "Bob"
	var isboy bool = true
	var prn int = 101010011
	var num float32 = 34.5345

	fmt.Printf("\nThe type are \n%v : %T\n%v : %T\n%v : %T\n%v : %T\n", name, name, isboy, isboy, prn, prn, num, num)

	// default values

	var sirname string
	fmt.Printf("\nThe default value for string is : %v\n", sirname)

	// declarations

	var college = "Engineering"
	fmt.Println(college)

	// no var style

	mobnum := 100100010
	fmt.Println(mobnum)
}
