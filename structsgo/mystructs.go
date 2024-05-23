package structsgo

import "fmt"

type engine interface {
	milesLeft() uint8
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
	int
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg

}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh

}
// func canMakeIt(e gasEngine, miles uint8) {
// 	if miles <= e.milesLeft() {
// 		fmt.Println("You can make it there")
// 	} else {
// 		fmt.Println("Need to fuel it first")
// 	}
// }

func canMakeIt(e engine, miles uint8) string{
	if miles <= e.milesLeft() {
		return "You can make it there"
	} else {
		return "Need to fuel it first"
	}
}
func Mystructs() {
	// 	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 200, ownerInfo: owner{name: "Michel"}}
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 200, owner: owner{name: "Michel"}, int: 10}
	var myEngine1 electricEngine = electricEngine{mpkwh: 25, kwh: 200, owner: owner{name: "Kim"}, int: 20}

	// anonymous struct not reusable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 105}

	fmt.Printf("\n mpg: %v gallons:%v \n", myEngine2.mpg, myEngine.gallons)

	fmt.Printf("\n Total miles left in tank: %v", myEngine.milesLeft())

	canMakeIt(myEngine1, 70)

	fmt.Printf("\n Can make it gasEngine: %v", canMakeIt(myEngine, 200))
	fmt.Printf("\n Can make it electricEngine: %v", canMakeIt(myEngine1, 200))

	// fmt.Printf("\n %v %v %v \n", myEngine.gallons, myEngine.mpg,myEngine.ownerInfo.name )
	fmt.Printf("\n %v %v %v %v \n", myEngine.gallons, myEngine.mpg, myEngine.name, myEngine.int)

}
