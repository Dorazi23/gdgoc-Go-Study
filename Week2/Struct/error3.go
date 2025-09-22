package main

import (
	"errors"
	"fmt"
)

var SampleError = errors.New("This is a test error")

func testPanic() {
	panic(SampleError)
	fmt.Println("Hello from testPanic!")
}

func testRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("got an error:", r)
		} else {
			fmt.Println("no error")
		}
	}()

	fmt.Println("Hello from testRecover!")
	testPanic()
	fmt.Println("This won't be printed")
}

func main() {
	testRecover()
}
