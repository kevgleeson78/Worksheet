//worksheet problem 10
//Reverse String
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet
//date 30/09/2017

package main

import (
	"fmt"
)

//Function that takes in and returns a String
//This ia the same as the reverse string used for the Palindrome test
//reused this section for this problem
func reverse(str string) string {
	//initialize two variables
	//to be able to manipulate them in a loop the rune[] is needed
	word := []rune(str)
	reverse := []rune{}
	// loop through the word and append the character
	//at the start of the reverse variable
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	//cast reverse back into a string
	newReverse := string(reverse)
	//return the reversed word
	return newReverse

}
func main() {
	//Input string to reverse and print out
	fmt.Println(reverse("kevin is here"))
}
