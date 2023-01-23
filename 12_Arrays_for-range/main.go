package main

import "fmt"

func main() {
	// Aynı veri tipinden bi sürü elemanı bir liste içinde tutmak için array kullanılır.
	/* city1 := "İstanbul"
	city2 := "Roma"
	city3 := "Milan"
	city4 := "Paris"

	fmt.Println(city1, city2, city3, city4) */

	//cities := [4]string{"İstanbul", "Roma", "Milan", "Paris"}
	/* cities := []string{"İstanbul", "Roma", "Milan", "Paris"}
	fmt.Println(cities)
	fmt.Println(cities[0])

	fmt.Println(len(cities)) */

	/* 	var myArr [5]int
	   	fmt.Println(myArr)

	   	myArr[0] = 5
	   	fmt.Println(myArr)

	   	myArr[len(myArr)-1] = 50
	   	fmt.Println(myArr) */

	/* 	var myArr [5]int
	   	var myArr2 [4]int
	   	if myArr == myArr2 {      // Bu iki arayın veri tipi birbirine eşit değil!!.
	   		fmt.Println("Mesaj")
	   	} */
	/* 	var myArr [4]int
	   	var myArr2 [4]int
	   	if myArr == myArr2 { // Bu iki arayın veri tipi birbirine eşit.
	   		fmt.Println("Mesaj")

	   	}*/
	/* 	cities := []string{"İstanbul", "Roma", "Milan", "Paris"}

	   	for i := 0; i < len(cities); i++ {
	   		fmt.Println(i, cities[i])

	   	} */
	/* myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(myArr)

	myArr = mySquare(myArr) //First class function

	fmt.Println(myArr) */

	//FOR -- RANGE yapisi
	// for index, value := range arr {}
	cities := []string{"İstanbul", "Roma", "Milan", "Paris"}
	for i, city := range cities {
		fmt.Println(i, city)
	}
}

/* func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]

	}
	return arr
} */
