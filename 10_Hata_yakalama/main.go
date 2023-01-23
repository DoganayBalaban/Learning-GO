/* package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Girdiğiniz sayı:", result)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA:Çift sayı girmediniz !!")

	}
	return num, nil //Başlangıç değeri olmayan ifadelere verilen değerdir.

} */

/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)

	} else {
		println("Girdiğiniz sayının karekökü:", result)
	}

}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Hata: Negatif sayının karekökünü alamazsınız.")

	}
	return math.Sqrt(num), nil
} */

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt") //Bu metotun kendi hata döndürmesi vardır hatayı yakalamak gerekmektedir.

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Dosyamız", file)
	}
}
