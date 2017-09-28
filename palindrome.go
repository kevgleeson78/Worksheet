package main

import (
	"fmt"
)

func palindrome(str string) string {
	word := []rune(str)
	reverse := []rune{}

	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	newReverse := string(reverse)
	newWord := string(word)

	if newReverse == newWord {
		return "the word you entered a palindrome"
	}
	return "the word you entered  is not a palindrome"

}
func main() {
	fmt.Println(palindrome("moooooom"))
}
