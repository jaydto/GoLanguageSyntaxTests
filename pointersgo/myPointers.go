package pointersgo

import "fmt"

func MyPointers() {
	var p *int32 = new(int32)
	var i int32
	var slice = []int{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 7

	var thing1 = [5]float64{1, 2, 3, 4}
	var result [5]float64 = square(&thing1)

	fmt.Println("square values", result)
	fmt.Printf("\n The pointer location of square values %p", &thing1)
	fmt.Printf("\n The pointer location of square values %p", &result)

	fmt.Printf("\n The value p points to is: %v \n ", *p)
	*p = 10
	fmt.Printf("\n The value p is now: %v \n ", *p)

	p = &i

	*p = 1

	fmt.Printf("\n The value of i is :%v \n ", i)

	fmt.Println("The value is: ", slice)
	fmt.Println("The value is: ", sliceCopy)

}

func square(things *[5]float64) [5]float64 {

	fmt.Printf("\n the memory location of thing 2 is:%p \n", &things)
	for i := range things {
		things[i] = things[i] * things[i]
	}
	return *things

}
