package interfacesgo

import "fmt"


type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meaow"
}


func Interfaces1(){
	var animal Animal
	animal = Dog{}
	fmt.Println(animal.Speak())

	animal = Cat{}
	fmt.Println(animal.Speak())
}
