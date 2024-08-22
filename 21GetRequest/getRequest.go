package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("WEb Requests")

	// fetchData()
	postData()
}

func fetchData() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Content length : ", res.ContentLength)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	resString.Write(content)

	// fmt.Println("Fetched Content : ", byteCount)
	fmt.Println("Fetched Content : ", resString.String())
	// fmt.Println("Fetched Content : ", string(content))
}

func postData() {
	const url = "https://jsonplaceholder.typicode.com/posts"

	data := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}
	reqBody, err := json.Marshal(data)

	res, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer(reqBody))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}

func postFormData() {
	const url = "fgfg"

	data := url.Values{}
	data.Add("name", "dixit")
	data.Add("id", strconv.Itoa(17))

	res, err := http.PostForm(url, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
