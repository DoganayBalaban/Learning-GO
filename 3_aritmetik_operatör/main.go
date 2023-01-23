package main

import "fmt"

func main() {
	x, y := 15, 10
	//fmt.Printf("%T, %v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("XT, %v\n", (x + y), (x + y))    toplama
	//fmt.Printf("XT, %v\n", (x - y), (x - y) )   çıkarma
	//fmt.Printf("NT, %v\n", (x * y), (x * y))    çarpma
	//fmt.Printf("%T, %v\n", (x / y), (x / y))    bölme
	fmt.Printf("%T, %v\n", (x % y), (x % y)) //bölünmünden kalan

	//z := 5.0 / 2 //(2.5)

	//fmt.Printf("%T, %v\n", z, z)

	a := 10

	a++ //1 arttır
	a-- //1 azalt
	fmt.Println(a)

}
