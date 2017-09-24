package main

import (
	"fmt"
)

func add(number int) int {

	sum := number % 10

	if number/10 < 10 {
		return sum + number/10
	} else {
		return sum + add(number/10)
	}
}
func main() {
	fmt.Println(add(123456))
}
