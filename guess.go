package main

import (
	"fmt"
	"math/rand"
	"time"
)

//function to accept the range of the number to randomly generate.
func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	//random number kept outside of loop to keep
	//one single random value
	rand.Seed(time.Now().UnixNano())
	randNum := random(1, 100)
	//foor loop with no arguments simular to while loop
	for {

		input := 0
		//guess := 0

		fmt.Println("Please enter a number")

		fmt.Scan(&input)

		//test input and output of user and random number.
		fmt.Println("The number entered is...", input, randNum)

		//conditionals for the user guess

		if input > randNum {
			fmt.Println("Youre guess is too high")
		} else if input < randNum {
			fmt.Println("Youre guess is too low")
		} else {
			fmt.Println("You have guesse the right number")
		}
		//break out of loop when the number is guessed correctly
		if input == randNum {
			break
		}
	}
}
