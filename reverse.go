package main

import (
	"fmt"
)

func reverse(str string) string {
	word := []rune(str)
	reverse := []rune{}

	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	newReverse := string(reverse)

	return newReverse

}
func main() {
	fmt.Println(reverse("kevin"))
}
