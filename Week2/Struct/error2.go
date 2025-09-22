// 패닉과 리커버
package main

import (
	"errors"
	"fmt"
)

var SampleError = errors.New("This is a test error")

func testRecover() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	panic(SampleError)
	fmt.Println("Hello!")
}

func main() {
	testRecover()
}
