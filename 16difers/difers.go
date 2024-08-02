package main

import "fmt"

func main() {
	fmt.Println("Wc")

	defer fmt.Println("At Last1")
	defer fmt.Println("At Last2")

	myDefer()
	fmt.Println("first")
	defer fmt.Println("At Last3")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
