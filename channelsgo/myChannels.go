package channelsgo

import (
	"fmt"
	"math/rand"
	"time"
)

func MyChannels() {
	// testChannelss()
	checkPrices()

}

func testChannelss() {
	var c = make(chan int, 5)
	go process(c)

	for i := range c {
		fmt.Printf("\nchannel %v\n", i)
		time.Sleep(time.Second + 1)
	}
}

func process(c chan int) {
	defer close(c)

	for i := 0; i < 5; i++ {
		c <- 1

	}

	fmt.Println("Exiting the procsess")
}

var MAX_CHICKEN_PRICE float32 = 5

func checkPrices() {

	var chickenPrice = make(chan string)
	var websites = []string{"walmart", "naivas", "tuskys"}

	for i := range websites {

		go checkChickenPrices(websites[i], chickenPrice)

	}
	sendMessage(chickenPrice)

}

func checkChickenPrices(website string, chickenPrice chan string) {

	for {
		time.Sleep(time.Second + 1)
		var chickenPricing = rand.Float32() * 20
		if chickenPricing <= MAX_CHICKEN_PRICE {
			chickenPrice <- website

		}
	}

}

func sendMessage(chickenPrice chan string) {

	fmt.Printf("\n Found a deal on chickens %v \n", <-chickenPrice)

}
