package main

import "fmt"

func main() {
	fmt.Println("Wc")

	cnt := 17
	var res string
	if cnt < 17 {
		fmt.Println("ok")
		res = "less"
	} else if cnt == 17 {
		res = "equal"
	} else {
		fmt.Println("more")
		res = "more"
	}

	if cnt = 25; cnt < 17 {
		fmt.Println("ok")
		res = "less"
	} else if cnt >= 17 {
		res = "equal"
	}

	fmt.Println(res)
}
