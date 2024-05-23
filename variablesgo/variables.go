package variablesgo

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

func VariablesTest() {

	// var age int
	var name string
	var marritalStatus bool
	const pi float64 = 3.1415
	age, age2 := 25, 30
	marritalStatus = false
	name = "Michek Angelo"
	fmt.Println(pi)

	fmt.Println("name:", name+" age: "+strconv.Itoa(age))
	fmt.Println(name+" age: ", age)
	fmt.Println(name+" age2: ", age2)
	fmt.Println("the time now is: ", time.Now())
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println("Marrital status", marritalStatus)

	printMe("print me")
	fmt.Println("summation", add(20, 24))
	var result1, result2, err = intDivision(24, 7)
	fmt.Println("value division", result1)
	fmt.Println("remainder division", result2)

	if err != nil {
		fmt.Println(err.Error())

	} else if result2 == 0 {
		fmt.Printf(" The result of this operation is :%v %v", result1, "\n")
	} else {
		fmt.Printf(" The value is : %v and remainder: %v of this division  %v", result1, result2, "\n")

	}

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case result2 == 0:
		fmt.Printf(" The results of this operation is : %v %v", result1, "\n")
	default:
		fmt.Printf("The value is %v and remainder is %v %v", result1, result2, "\n")
	}

	switch result2 {
	case 0:
		fmt.Println(" The remainder is :", 0)
	case 1, 2:
		fmt.Printf("the remainder is either 1 or 2 currently its : %v %v ", result2, "\n")
	default:
		fmt.Printf(" This is a switch clause value is: %v %v", result2, "\n")
	}

	mod(23, 7)
	division(36, 5)

	

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func add(num1 int, num2 int) int {
	value := num1 + num2
	return value
}

func mod(num1 int, num2 int) {
	value := num1 % num2
	println(value)
}

func intDivision(num1 int, num2 int) (int, int, error) {

	var err error

	if num2 == 0 {
		err = errors.New("Cannot Divide by Zero")

		return 0, 0, err
	}

	var value = num1 / num2

	var remainder = num1 % num2

	return value, remainder, err

	// printMe(strconv.FormatFloat(value, 'f', 2, 64))

}
func division(num1 int, num2 int) {

	var value float64 = float64(num1) / float64(num2)

	printMe(strconv.FormatFloat(value, 'f', 2, 64))

}

