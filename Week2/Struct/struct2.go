// 구조체의 올바른 참조
package main

import (
	"fmt"
)

func incrementAge(user *User) {
	user.Age++
	fmt.Println(user.Age)
}

// 올바른 참조
func main() {
	dorazi := User{"dorazi", 19}
	incrementAge(&dorazi)
	fmt.Println(dorazi.Age)
}
