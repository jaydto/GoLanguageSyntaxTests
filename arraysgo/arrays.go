package arraysgo

import "fmt"

func ArrayTests() {

	// var numbers [3]int32 = [3]int32{1,2,3}
	numbers := [3]int32{1, 2, 3}

	// slices

	var numbers2 []int32 = make([]int32, 8, 9)



	fmt.Printf("\n The length of this numbers is %v and its capacity : %v of this make option \n", len(numbers2), cap(numbers2))

	stringvalues := []string{"name1", "name2", "name3"}

	stringvalues3 := []string{"name5", "name6", "name7"}

	fmt.Printf("\n  The length is %v and the capacity is: %v before appending  \n", len(stringvalues), cap(stringvalues))

	stringvalues = append(stringvalues, "name4")
	fmt.Println(stringvalues)

	stringvalues = append(stringvalues, stringvalues3...)

	fmt.Printf("\n  The length is %v and the capacity is %v after appending \n", len(stringvalues), cap(stringvalues))

	fmt.Println(stringvalues)

	fmt.Println(numbers[0])

	fmt.Println(numbers[0:3])

	fmt.Println(&numbers[0])

}
