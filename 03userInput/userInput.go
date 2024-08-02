package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomMsg := "Welcome to First Lab"
	fmt.Println(welcomMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give feedback point : ")

	// commac ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("THanks for ", input)
	fmt.Printf("Type of feedback  : %T", input)
}
