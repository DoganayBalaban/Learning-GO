package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Seafood", "foo")) //foo kelimesi seafoodun içinde var mı

	fmt.Println(strings.Count("Merhaba", "a")) //Merhabanın içinde kaç tane a var.
}
