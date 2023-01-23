package main

import "fmt"

func main() {
	// if <bool cümle> {kod}  bool cümle false ise kod satırına girmez
	/* x := 27
	if x%2 == 0 {
		fmt.Println(x, "çift sayidir")
	} else {
		fmt.Println(x, "tek sayidir")
	}
	if x%2 != 0 {
		fmt.Println(x, "çift sayi degildir.")

	}*/

	/* x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayidir")     // 1. doğru olduktan sonra hiç birine bakmadı.

	} else if x%2 != 0 {
		fmt.Println(x, "Tek sayidir")

	} else {
		fmt.Println(x, "Yanlis ifade")
	} */

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayidir") // 1. doğru olduktan sonra hiç birine bakmadı.

	} else if x%2 != 0 {
		fmt.Println(x, "Tek sayidir")

	} else {
		fmt.Println(x, "Yanlis ifade")
	}

}
