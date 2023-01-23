package main

import "fmt"

func main() {
	var name string
	name = "Doganay"

	name2 := "Emir"

	var age int = 19

	var isStudent bool
	isStudent = true

	var weight = 71.5

	/* var (
	 	school   = "KLU"
		sinif    = 1
		ortalama = 80.2
	 ) */

	/* var school, sinif, ortalama = "KLU", 1, 80.2 */

	school, sinif, ortalama := "KLU", 1, 80.2

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(weight)
	fmt.Println(school)
	fmt.Println(sinif)
	fmt.Println(ortalama)

	fmt.Printf("\n%T\n", name)
	fmt.Printf("%T\n", name2)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isStudent)
	fmt.Printf("%T\n", weight)
}
