package main

import "fmt"

type User struct {
	Name  string
	Score int
	Live  bool
}

func main() {
	fmt.Println("Wc")

	dixit := User{"ld", 17, true}
	fmt.Println(dixit)
	fmt.Printf("dixit : %+v\n", dixit)
	fmt.Printf("dixit : %v\n", dixit.Name)

	dixit.getStatus()
}

func (u User) getStatus() {
	fmt.Println("is alive", u.Live)
}
