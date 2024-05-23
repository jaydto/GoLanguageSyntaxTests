package genericsgo

import (
	"encoding/json"
	"fmt"
	"log"

	// "io"
	"io/ioutil"
	// "path/filepath"
)

type contactInfo struct {
	name  string
	email string
}

type purchaseInfo struct {
	name   string
	price  float32
	smount int
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func carOptions() {
	// var gasCar = car[gasEngine]{carMake: "Honda", carModel: "2018", engine: gasEngine{mpg: 20, gallons: 50}}
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func loadJson1() {
	var contacts []contactInfo = loadJSON[contactInfo]("contactInfo.json")
	fmt.Printf("\n %v\n", contacts)

	var purchase []purchaseInfo = loadJSON[purchaseInfo]("purchaseInfo.json")
	fmt.Printf("\n %v\n", purchase)
}

func loadJSON[T any](filepath string) []T {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var loaded []T
	if err := json.Unmarshal(data, &loaded); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return loaded
}

func MyGenerics() {

	var instSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumInt[int](instSlice))

	var floatslice = []float64{1.2, 3.4, 4.5}
	fmt.Println(sumInt[float64](floatslice))
	fmt.Println(isEmpty[float64](floatslice))

	loadJson1()

}

func sumInt[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum

}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0

}
