package runesgo

import (
	"fmt"
	"unicode/utf8"
)

func RunesGo() {

	myString := []rune("r´e™😂sume")
	myString2 := "r´e™😂sume"

	var myString3 = []rune("résumé")

	indexedString := myString[1]

	// fmt.Println(myString)
	// fmt.Println(len(myString))
	fmt.Println(len(myString2))
	fmt.Println(len(myString3))

	// fmt.Println(utf8.RuneLen(myString3))

	fmt.Println(utf8.RuneCountInString(myString2))
	fmt.Printf("\n indexed string and value %v %T", indexedString, indexedString)

	// for i, v := range myString {

	// 	fmt.Printf("\n indexed values: value, position and type %v %v %T ", i, v, v)

	// }

}
