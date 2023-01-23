package main

import "fmt"

func main() {
	/* x, y := 3, 7

	fmt.Printf("%v\n", x == y)
	fmt.Printf("%v\n", x != y)
	fmt.Printf("%v\n", x < y)
	fmt.Printf("%v\n", x > y)
	fmt.Printf("%v\n", x >= y)
	fmt.Printf("%v\n", x <= y) */

	x, y := 3, 5

	set1 := (x == y) //false
	set2 := (x < y)  //true
	set3 := true

	fmt.Printf("%v\n", (set1 && set2)) //0 and 1 --> 0
	fmt.Printf("%v\n", (set1 || set2)) //0 or 1 --> 1
	fmt.Printf("%v\n", (!set3))

}
