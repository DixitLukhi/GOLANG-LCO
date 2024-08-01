package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.iana.org/assignments/media-types/application/vnd.api+json"

func main() {
	fmt.Println("Wc")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Respinse type %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data : ", string(databytes))
}
