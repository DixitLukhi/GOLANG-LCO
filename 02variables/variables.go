package main

import "fmt"

// users := 450 // Not allowed outside function
var users int = 450

// First char is capital means public - not mandatory
const LoginPass string = "abc123"

func main() {
	fmt.Println("Hi")

	var username string = "Dixit"
	fmt.Println(username)
	fmt.Printf("Type : %T ", username)

	var status bool = true
	fmt.Println(status)
	fmt.Printf("Type : %T ", status)

	var val int = 17
	fmt.Println(val)
	fmt.Printf("Type : %T ", val)

	var floatVal float32 = 17.93281932819328193281
	fmt.Println(floatVal)
	fmt.Printf("Type : %T ", floatVal)

	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Type : %T ", defaultVal)

	totalUsers := 5000
	fmt.Println(totalUsers)

	fmt.Println(users)

	fmt.Println(LoginPass)
}
