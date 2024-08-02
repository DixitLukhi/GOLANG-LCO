package main

import "fmt"

func main() {
	fmt.Println("Wc")

	langs := make(map[string]string)

	langs["JS"] = "JavaScript"
	langs["C"] = "C"
	langs["Cpp"] = "C++"
	langs["PY"] = "Python"

	fmt.Println("Langs : ", langs)
	fmt.Println("Langs : ", langs["PY"])

	delete(langs, "Cpp")

	fmt.Println("Langs : ", langs)

	for _, value := range langs {
		// fmt.Printf("For %v, value is %v\n", key, value)
		fmt.Println("For", "=", value)
	}
}
