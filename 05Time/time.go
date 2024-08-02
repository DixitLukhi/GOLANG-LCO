package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to random no : ")

	// math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// crypto/rand
	randomNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNo)

	fmt.Println("Welcome to time : ")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, -1, 24, 24, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
