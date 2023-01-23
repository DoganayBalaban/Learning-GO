package main

import "fmt"

func main() {
	// Arrayler oluşturulduğunda kendi eleman sayısını bilirken sliceslar oluşturulduğunda kendi eleman sayısını bilmez.
	/* mySlc := [4]int{1, 2, 3, 4}
	fmt.Println(mySlc)
	fmt.Println(len(mySlc)) */

	/* 	var mySlc []int
	   	fmt.Println(mySlc) */

	/* var myArr [4]int
	fmt.Println(myArr)
	myArr[0] = 5
	fmt.Println(myArr)

	var mySlc []int
	mySlc = make([]int, 4)
	fmt.Println(mySlc)
	mySlc[0] = 10
	fmt.Println(mySlc) */

	/* myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)

	myArr2[0] = 100          // Diziler sadece değerlerini paylaşırken
	fmt.Println(myArr2)
	fmt.Println(myArr) */

	/* 	mySlc := []int{1, 2, 3}
	   	fmt.Println(mySlc)
	   	mySlc2 := mySlc
	   	fmt.Println(mySlc2)
	   	// Sliceler referanslarını da paylaşır.
	   	mySlc2[0] = 100
	   	fmt.Println(mySlc2)
	   	fmt.Println(mySlc) */

	/* 	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	   	fmt.Println(underArray)

	   	mySlc := underArray[2:5]
	   	fmt.Println(mySlc)
	   	mySlc2 := underArray[:6]
	   	fmt.Println(mySlc2)
	   	mySlc3 := underArray[3:]
	   	fmt.Println(mySlc3) */

	/* mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	mySlc = append(mySlc, 4, 5)
	fmt.Println(mySlc) */
	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5, 6}
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc) */

	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlc)
	mySlc = mySlc[:len(mySlc)-3]
	fmt.Println(mySlc)

}
