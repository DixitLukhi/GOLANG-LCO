package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wc")

	var nameList = []string{"dixit", "Neha"}
	fmt.Printf("Type : %T\n", nameList)
	fmt.Println("List : ", nameList)

	nameList = append(nameList, "Wooho")
	fmt.Println("List : ", nameList)
	nameList = append(nameList[0:1])
	fmt.Println("List : ", nameList)

	scores := make([]int, 5)

	scores[0] = 1
	scores[1] = 14
	scores[2] = 4540
	scores[3] = 031
	scores[4] = 48450

	scores = append(scores, 4545, 78878)
	fmt.Println("scores : ", scores)

	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	var ind int = 2
	scores = append(scores[:ind], scores[ind+3:]...)
	fmt.Println("remove : ", scores)

}
