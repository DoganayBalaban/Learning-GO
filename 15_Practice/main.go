// 1 -) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.
/* package main

import "fmt"

func main() {
	myArr := [5]string{"Fenerbahçe", "Galatasaray", "Beşiktaş", "Trabzonspor", "Başakşehir"}
	for index, value := range myArr {
		fmt.Println(index, value)
	}
	fmt.Println()
	mySlc := []int{29, 25, 19, 5, 1}
	for i, v := range mySlc {
		fmt.Println(i, v)
	}

} */

// 2 -) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluşturunuz.
/* package main

import "fmt"

func main() {
	mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mySlc)
	mySlc2 := mySlc[0:4]
	fmt.Println(mySlc2)
	mySlc3 := mySlc[4:7]
	fmt.Println(mySlc3)
	mySlc4 := append(mySlc[2:4], mySlc[6:8]...)
	fmt.Println(mySlc4)

} */

// 3 -) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.
/* package main

import "fmt"

func main() {

} */
// 4 -) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.
package main

import "fmt"

func main() {
	kitapYazar := map[string]string{
		"Dostoyevski":           "Suç ve Ceza",
		"Jack London":           "Martin Eden",
		"Mustafa Kemal Atatürk": "Nutuk",
		"Franz Kafka":           "Dönüşüm",
		"George Orwell":         "1984",
	}
	for key, values := range kitapYazar {
		fmt.Println("Yazar: ", key)
		fmt.Println("Eser: ", values)
		fmt.Printf("\n\n")
	}
}
