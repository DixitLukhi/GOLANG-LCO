package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Id      string
	Name    string `json:"username"`
	Student bool
	Faltu   string   `json:"-"`
	About   []string `json:"about,omitempty"`
}

func main() {
	fmt.Println("JSON")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	users := []info{
		{"17", "dixit", true, "naaaa", []string{"good"}},
		{"14", "shrut", true, "passworddd", []string{"good2"}},
		{"20", "lukhi", false, "hahhaha", nil},
	}

	finalJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}

	// fmt.Println(finalJson)
	fmt.Printf("%s\n", finalJson)
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
