package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang !!!")
	text := "Hello there, I wrote this in file using a program !!!"

	file, fileerr := os.Create("./file.txt")
	checkerror(fileerr)

	length, writeerr := io.WriteString(file, text)
	checkerror(writeerr)
	fmt.Println("The length of the file is ", length)
	readFile("./file.txt")
	defer file.Close()
}

func readFile(filepath string) {
	databye, err := ioutil.ReadFile(filepath)
	checkerror(err)
	fmt.Println("The data in file is : ", string(databye))
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
