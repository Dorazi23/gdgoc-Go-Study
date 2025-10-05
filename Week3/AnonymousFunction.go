package main

import (
	"fmt"
)

func main() {
	inputChan := make(chan int, 10)
	finishChan := make(chan int)
	outputChan := make(chan int, 10)
	go func(inputChan, finishChan chan int) {
		for {
			select {
			case x := <-inputChan:
				outputChan <- x * x
			case _ = <-finishChan:
				return
			}
		}
	}(inputChan, finishChan)

	for i := 0; i < 10; i++ {
		inputChan <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChan)
	}
	finishChan <- 1
}
