package main

import "fmt"

func greet() {
	fmt.Println("from greet")
}

func main() {
	greet()
	fmt.Println("Wc")

	res := sum(5, 2)
	fmt.Println("add : ", res)

	moreSum, str := proAdder(5, 1, 2, 0, 012)
	fmt.Println(moreSum, str)
}

func sum(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(vaues ...int) (int, string) {
	total := 0

	for _, val := range vaues {
		total += val
	}
	return total, "From sec args"
}
