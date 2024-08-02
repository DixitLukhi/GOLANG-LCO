package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	var namelist [3]string
	namelist[0] = "dixit"
	namelist[1] = "lukhi"

	fmt.Println("names : ", namelist)
	namelist[2] = "manan"
	fmt.Println("names : ", namelist)
	fmt.Println("names : ", len(namelist))

	var surnameList = [5]string{"lukhi", "raval", "dix", "LS"}
	fmt.Println("names : ", surnameList)
	fmt.Println("names : ", len(surnameList))
}
