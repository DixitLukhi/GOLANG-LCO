package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Id      string
	Name    string
	Student bool
}

func main() {
	fmt.Println("json")

	decodeJson()
}

func decodeJson() {
	jsonData := []byte(`
	{
		"Id": "17",
		"Name": "dixit",
		"Student": true
	}`)

	var users info
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON VALID")
		json.Unmarshal(jsonData, &users)
		fmt.Printf("%#v", users)
	} else {
		fmt.Println("JSON data not valid")
	}

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("key : %v, value : %v, type : %T\n", k, v, v)
	}
}
