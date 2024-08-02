package main

import "fmt"

func main() {
	fmt.Println("Hi")

	var ptr *int
	fmt.Println("Ptr : ", ptr)

	myNum := 17

	var ptr1 = &myNum
	fmt.Println("mynum address: ", &myNum)
	fmt.Println("Ptr1 address: ", ptr1)
	fmt.Println("Ptr1 : ", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("New val : ", myNum)
}
