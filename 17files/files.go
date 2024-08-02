package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Wc")

	content := "This will be in the file"

	file, err := os.Create("./temp.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./temp.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)

	if err != nil {
		panic(err)
	}

	fmt.Println("In file : ", string(databyte))
}
