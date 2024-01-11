package main

import (
	"learning-go/helpers"
	"log"
)

const numPool = 10

func ProcessValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool) // generate value
	intChan <- randomNumber                       // push random value into channel
}

func main() {
	// Creating a channel, recieves values to hold, in this case integers
	intChan := make(chan int)
	defer close(intChan)

	go ProcessValue(intChan) // Create a goroutine with the channel

	num := <-intChan // Receive the value from the channel
	log.Println(num)
}
