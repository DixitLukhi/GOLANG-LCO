package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomMsg := "Welcome to First Lab"
	fmt.Println(welcomMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give feedback point : ")

	// commac ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("THanks for ", input)

	numRate, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add : ", numRate+1)
	}
	fmt.Printf("Type of feedback  : %T", input)
}
