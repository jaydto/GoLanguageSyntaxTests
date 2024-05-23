package stringsgo

import (
	"fmt"
	"strings"
)

func StringManipulations() {

	var strSlice = []string{"s", "a", "m"}

	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	fmt.Println("\n hy",strBuilder.String())

	castr := ""

	for i := range strSlice {

		castr += strSlice[i]

	}

	fmt.Printf("\n hello %v \n ", castr)

}
