package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//fmt.Println(x + y) farklı veri tiplerini toplayamaz.

	// Type Conversion - type(value) => int(y) = 10.0 => 10
	fmt.Println(x + int(y))

	fmt.Printf("%v %T\n", y, y)

	num1 := 106
	str1 := string(num1)

	fmt.Printf("\n%v %T\n", num1, num1)
	fmt.Printf("\n%v %T\n", str1, str1) //106 j ifadesine denk geldiği için int -> str dönüşümünde j ye dönüşür.

}
