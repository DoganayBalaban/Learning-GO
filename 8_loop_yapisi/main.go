package main

import "fmt"

func main() {
	/* for i := 0; i <= 10; i++ {
		fmt.Println(i)

	}*/
	/*
		i := 0
		for ; i <= 5; i++ {
			fmt.Println(i)

		}
		fmt.Println(i) */

	/* for {           //Sonsuz döngü
		println("Merhaba")

	} */
	/* for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue // Burası sağlandığı anda ileri gitme for başına geri dön demek.

		}
		fmt.Println(i)

	} */

	for i := 0; i <= 10; i++ {
		if i == 3 {
			break // Burası sağlandığı anda loopu bitir.

		}
		fmt.Println(i)

	}

}
