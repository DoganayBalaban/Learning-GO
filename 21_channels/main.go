/*
	package main

import (

	"fmt"
	"math"

)

	type circle struct {
		r float64
	}

	func (c circle) display() {
		fmt.Println("A circle")
	}

	func (c circle) area() float64 {
		return math.Pi * c.r * c.r
	}

	func main() {
		c1 := circle{5}
		c1.display()
		area1 := c1.area()
		fmt.Printf("%.2f\n", area1)
	}
*/
/* package main

import "fmt"

func merhaba(chan1 chan string) {
	chan1 <- "Merhaba"
}

func main() {
	  var myChannel chan string
	myChannel = make(chan string)

	myChannel := make(chan string)

	go merhaba(myChannel)
	fmt.Println(<-myChannel)
} */
package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A circle")
}

/* func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}*/

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}
func main() {
	c1 := circle{5}
	chan1 := make(chan float64)
	go area(c1, chan1)
	fmt.Printf("%.2f\n", <-chan1)
	c1.display()
}
