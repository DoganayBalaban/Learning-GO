/*
	package main

import "fmt"

	func main() {
		var öğrenciDurum struct {
			name     string
			exam     int
			isPassed bool
		}
		fmt.Println(öğrenciDurum)
		fmt.Println(öğrenciDurum.exam)
		öğrenciDurum.name = "Doğanay"
		öğrenciDurum.exam = 85
		öğrenciDurum.isPassed = true
		fmt.Println(öğrenciDurum)
	}
*/
/* package main

import "fmt"

type öğrenciDurum struct {
	name     string
	exam     int
	isPassed bool
}

func main() {

	var o1 öğrenciDurum
	o1.name = "Doğanay"
	o1.exam = 85
	o1.isPassed = true

	var o2 öğrenciDurum
	o2.name = "Eda"
	o2.exam = 35
	o2.isPassed = false

	fmt.Println(o1)
	fmt.Println(o2)
} */
package main

import "fmt"

type öğrenciDurum struct {
	name     string
	exam     int
	isPassed bool
	lessons  []string
}
type schoolManager struct {
	öğrenciDurum
	isDegree bool
}

func main() { //pass by value
	/* e1 := öğrenciDurum{
		name:     "Doğanay",
		exam:     85,
		isPassed: true,
		lessons: []string{
			"Math",
			"Algorithm",
		},
	}
	fmt.Println(e1)
	fmt.Printf("%#v\n", e1) */

	m1 := schoolManager{
		öğrenciDurum: öğrenciDurum{
			name:     "Eda",
			exam:     35,
			isPassed: false,
			lessons: []string{
				"Math",
				"Pyshics",
			},
		},
		isDegree: false,
	}
	fmt.Println(m1)

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}
	fmt.Println(theBoss)
}
