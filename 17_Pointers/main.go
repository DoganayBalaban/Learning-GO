package main

import "fmt"

func main() {
	/*x := 20

		fmt.Println(x)
	  	fmt.Println(&x) //--> Adres operatörü

	  	fmt.Printf("%T,%v\n", x, x)
	  	fmt.Printf("%T,%v\n", &x, &x) */
	/* 	x := 20
	   	/* y := &x
	   	fmt.Printf("%T,%v", y, y)
	   	fmt.Println(x)     //20
	   	fmt.Println(&x)    //adress
	   	fmt.Println(*(&x)) //20---dereferencing */

	/* x1 := 10
	x2 := x1
	fmt.Println(x1, x2)
	x1 = 5
	fmt.Println(x1, x2) */
	/* x1 := 10
	x2 := &x1
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)
	*x2 = 3
	fmt.Println(x1, *x2) */
	x1 := [4]int{1, 10, 100, 1000}
	x2 := x1
	fmt.Println(x1, x2)
	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1)
	//array = pass by value
	//slice = pass by referance
}
