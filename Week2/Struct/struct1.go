// 구조체의 형태
package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Dorazi", 24}
	case 1:
		return User{"Gorani", 16}
	default:
		return User{"", -1}
	}
}

func main() {
	user := nameAndAge(1)
	fmt.Println("user age:")
	fmt.Println(user.Age)
}
