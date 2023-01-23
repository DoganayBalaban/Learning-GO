package main

import "fmt"

func main() {
	grade := 3

	switch grade {
	case 5: // grade 5'e eşit mi ??
		fmt.Println("Başarili")

	case 4:
		fmt.Println("iyi")

	case 3:
		fmt.Println("orta") //Burdan gerisi çalışmaz.

	case 2:
		fmt.Println("gecer")

	case 1:
		fmt.Println("başarisiz")

	default: //else
		fmt.Println("Geçersiz not")
	}

	switch {
	case false:
		println("görünmez")
	case true:
		fmt.Println("görünür")
	}

}
