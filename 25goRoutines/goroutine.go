package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

var signals = []string{"test"}

var mut sync.Mutex

func main() {

	// go greeter("Hello")
	// greeter("World")

	wesiteList := []string{
		"https://lco.dev",
		"https://google.dev",
		"https://go.dev",
		"https://github.dev",
		"https://applogdaiict.netlify.app/",
	}

	for _, web := range wesiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("signals : ", signals)
}

func greeter(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Wrong endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
	}
}
