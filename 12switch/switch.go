package main

import "fmt"

func main() {
	fmt.Println("Wc")

	randNo := 8

	switch randNo {
	case 1:
		fmt.Println("Its 1")
	case 2:
		fmt.Println("Is 2")
		fallthrough
	case 5:
		fmt.Println("Its 5")
	case 6:
		fmt.Println("Its 6")
		fallthrough
	default:
		fmt.Println("Oops")
	}
}
