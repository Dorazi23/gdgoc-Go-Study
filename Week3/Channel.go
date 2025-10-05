package main

import (
	"fmt"
)

func squareIt(inputChan, outputChan chan int) {
	for x := range inputChan {
		outputChan <- x * x
	}
}

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10)

	go squareIt(inputChannel, outputChannel)

	for i := 0; i < 10; i++ {
		inputChannel <- i
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChannel)
	}
}
