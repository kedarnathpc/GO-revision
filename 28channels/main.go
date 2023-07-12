package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang !!!")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 9
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
