package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in golang !!!")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("pushing one")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("pushing two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("pushing three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
