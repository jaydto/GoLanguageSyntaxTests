package mapsgo

import "fmt"

func MyMaps() {
	var myMaps map[string]uint8 = make(map[string]uint8)

	myMaps2 := map[string]string{"name": "John", "age": "23", "specialization": "Software"}

	fmt.Printf("\nMaps information is %v \n", myMaps2)

	myMaps["main"] = 23

	value, ok := myMaps["now"]

	switch ok {
	case true:
		fmt.Printf("\n Value is : %v \n", value)
	default:
		fmt.Printf("\n The item is not available")
	}

	fmt.Printf("\n Maps options  %v \n", myMaps["main"])

	for name, value := range myMaps2 {
		fmt.Printf("\nMap items key: %v value: %v\n", name, value)
	}
	i := 0
	for i < 10 {

		// second option
		// if i>=10{

		// 	break
		// }

		// first option
		fmt.Println("value in while loop: ", i)
		i++

	}

	for i:=0;i<10;i++{
		fmt.Println("I am doing a secondary loop", i)
	}
}
