package main

import (
	"fmt"
)

func squareIt(x int) {
	fmt.Println(x * x)
}

func main() {
	go squareIt(2)
}
