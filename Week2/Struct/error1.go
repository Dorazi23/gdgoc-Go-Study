// 에러 인터페이스
package main

import (
	"errors"
	"fmt"
)

var DivisionByZero = errors.New("division by zero")

func Divide(number, d float32) (float32, error) {
	if d == 0 {
		return 0, DivisionByZero
	}
	return number / d, nil
}

func main() {
	n1, e1 := Divide(1, 1)
	fmt.Println(n1)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	n2, e2 := Divide(1, 0)
	fmt.Println(n2)
	if e2 != nil {
		fmt.Println(e2.Error())
	}
}
