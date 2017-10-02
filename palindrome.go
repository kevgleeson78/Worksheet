//worksheet problem 7
//@Author Kevin Gleeson
//date 19/09/2017

package main

import (
	"fmt"
)

//create function that takes in and returns a string
func palindrome(str string) string {
	//initialize two variables
	//to be able to manipulate them in a loop the rune[] is needed
	word := []rune(str)
	reverse := []rune{}
	// loop through the word and append the character
	//at the start of the reverse variable
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	//cast the word and reverse back into a string
	newReverse := string(reverse)
	newWord := string(word)
	//Conditional to check for equality of the strings
	//if they are the same the word is a palindrome
	if newReverse == newWord {
		return "the word you entered is a palindrome"
	}
	//If condition is not met return the below message
	return "the word you entered is not a palindrome"

}
func main() {
	//Input word for test
	fmt.Println(palindrome("moooooom"))
}
