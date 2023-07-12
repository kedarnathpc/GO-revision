package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	var signals = []string{"test"}
	websitelist := []string{
		"https://lco.dev",
		"https://amazon.com",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(&wg, web, &mut, &signals)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(wg *sync.WaitGroup, endpoint string, mut *sync.Mutex, signals *[]string) {
	defer wg.Done()
	response, err := http.Get(endpoint)
	checkError(err)

	mut.Lock()
	*signals = append(*signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
