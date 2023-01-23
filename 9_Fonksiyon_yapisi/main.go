/* package main

import "fmt"

func main() { */
/* 	var x, y, sum int
   	x = 5
   	y = 10
   	sum = x + y
   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)
   	x = 7
   	y = 11
   	sum = x + y
   	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum) */

// Moduler Programming
// Readable code
// From Complex To Simple

/* 	fmt.Println(sum(5, 10)) */

// func funcName(params) return type { code  }

/* 	merhaba()
   	fmt.Println(sum(5, 10))
   	fmt.Println(sum(3, 5))
   	fmt.Println(sum(2, 7))
   	merhaba()
	   merhaba() */

// fmt.Println(x)

// return vs print

/* 	z := sum(5, 10)
   	fmt.Println(z)
	   sum2(6, 11) */

/* merhaba2("Doğanay", 19) */ //Argüman fonksiyon çalıştırırken

// Fonksiyon İsimlendirme
// ilk karekter harf
// camel Case -- mySum, myBestFunction
// paket dışı --> ilk harf büyük

/* fmt.Println(result(55))
	fmt.Println(result(45))

}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim Adım Arin")
}

func merhaba2(name string, age int) { //Parametre fonksiyon yazarken
	fmt.Printf("\nAdım %s, yaşım %d\n", name, age)
}

func result(grade int) string {
	if grade >= 50 {
		return "Passed"
	}
	return "Not Passed"  // Dönüşlü fonksiyonlarda if kullanıyorsanız else kalıbına gerek yoktur.
}*/

/*
	package main

import (

	"fmt"
	"time"

)

	func main() {
		var x int = 10
		fmt.Println(x)

		var moment time.Time = time.Now() //metot
		fmt.Println(moment)

}
*/
/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Sınav sonucunuzu girin: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // bu metot geriye iki sonuç döner birini value ile yakalarız ikincisi ise errordur ama  _ --> blank identifier ile yakaladık.
	fmt.Println(value)                  // _ yani blank identifier kullanırsan o veriyi kullanmana gerek kalmaz.

} */

package main

import "fmt"

func main() {
	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm)
	fmt.Println(kalan)

}

func bölme(bölünen, bölen int) (kalan, bölüm int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen
	return bölüm, kalan

}
