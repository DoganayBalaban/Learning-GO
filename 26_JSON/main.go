/*
******JSON Marshal*****
 */

/* package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type kişi struct {
	İsim    string
	Soyisim string
	Yaş     int
} */

/*
	 	Doğanay := kişi{
		İsim:    "Doğanay",
		Soyisim: "Balaban",
		Yaş:     19,
	}

veri, err := json.Marshal(Doğanay)

	if err != nil {
		log.Fatalln(err)
		return
	}

fmt.Printf("JSON Parse Sonucu:%s", string(veri))
*/
/* func main() {
	ali := kişi{
		İsim:    "Ali",
		Soyisim: "Veli",
		Yaş:     20,
	}
	veri, err := json.MarshalIndent(ali, "", "    ")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("JSON Parse Sonucu:\n%s", string(veri))
} */

/*
*********JSON Unmarshal******
 */
/* package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type kişi struct {
	İsim    string
	Soyisim string
	Yaş     int
}

func main() {
	jsonVeri := []byte(`{"İsim":"Latif","Soyisim":"Uluman","Yaş":23}`)
	var goVeri kişi
	err := json.Unmarshal(jsonVeri, &goVeri)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("İsim - Soyisim: %s %s\nYaş: %d", goVeri.İsim, goVeri.Soyisim, goVeri.Yaş)
} */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonVeri := []byte(`{"İsim":"Latif","Soyisim":"Uluman" ,"Yas":23 , "Kilo":80.25}`)
	var goVeri map[string]interface{}
	err := json.Unmarshal(jsonVeri, &goVeri)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return
	}
	fmt.Printf("İsim: %+v \nSoyisim: %+v \nYas:%+v\nKilo:%+v", goVeri["İsim"], goVeri["Soyisim"], goVeri["Yas"], goVeri["Kilo"])
}
