package main

import "fmt"

func main() {
	/* myMap := map[string]int{
		"Doğanay": 19,
		"Eda":     20,
		"Elif":    18,
		"Emir":    13,
	}
	fmt.Println(myMap)
	//fmt.Println(myMap[0]) yazamayız haya verir
	fmt.Println(myMap["Doğanay"]) */
	/* 	myMap := map[string]int{
	   		"Doğanay": 19,
	   		"Eda":     20,
	   		"Elif":    18,
	   		"Emir":    13,
	   	}
	   	fmt.Println(myMap["Ayşe"]) // map içinde olmayan değer zero values döner. */
	/* 	isMarried := map[string]bool{
		"Ahmet":  true,
		"Ayşe":   false,
		"Mahmut": false,
	} */
	/* fmt.Println(isMarried) */
	/* myMap := map[string]int{
		"Doğanay": 19,
		"Eda":     20,
		"Elif":    18,
		"Emir":    13,
	} */

	/* value, ok := myMap["Elis"]  //Elis değeri map içinde var mı yok mu belirteci.
	fmt.Println(value, ok) */

	/* _, ok := myMap["Elis"]
	fmt.Println(ok) */

	/* _, ok := myMap["Doğanay"]
	fmt.Println(ok) */

	myMap := map[string]int{
		"Doğanay": 19,
		"Eda":     20,
		"Elif":    18,
		"Emir":    13,
	}

	fmt.Println(myMap)
	myMap["Mahmut"] = 30
	fmt.Println(myMap)

	delete(myMap, "Elif")
	fmt.Println(myMap)
	fmt.Println(len(myMap))
}
