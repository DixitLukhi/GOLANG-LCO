package main

import "fmt"

func main() {
	fmt.Println("Wc")

	days := []string{"Mon", "Tue", "Wed", "Thur",
		"Fri"}

	fmt.Println(days)

	//  NO ++i
	for i := 0; i < len(days); i += 2 {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println("Index : ", index, "=", day)
	}

	num := 1
	for num < 11 {
		if num == 3 {
			goto label
		}
		if num == 5 {
			num++
			continue
		}
		if num == 7 {
			break
		}
		fmt.Println(num)
		num += 1
	}
label:
	fmt.Println("from label")

	fmt.Println("End")
}
