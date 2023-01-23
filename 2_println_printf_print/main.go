package main

import "fmt"

func main() {
	// print ve println olduğu gibi ham şekilde yazdırırken printf formatlanmış şekilde yazdırır.
	//ln yeni bir satır başlangıcı yapar
	/* fmt.Println("Merhaba")
	/* fmt.Print("Merhaba")
	/* fmt.Printf("Merhaba") */

	name := "Doganay"

	fmt.Print("Benim Adim", name)
	fmt.Println("")
	fmt.Println("Benim Adim", name)
	fmt.Printf("Benim Adim %v", name)

}
