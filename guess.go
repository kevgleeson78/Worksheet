package main

import (
	"fmt"
	"math/rand"
	"time"
)

//function to accept the min max range of the number to randomly generate.
func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	//random number kept outside of loop to keep
	//one single random value
	rand.Seed(time.Now().UnixNano())
	randNum := random(1, 100)
	//i variable to hold amount of attempts
	i := 1
	//input variable to hold the user input
	input := 0
	// sameNumber variable to hold the previous number to test if the user has entered the same number
	sameNumber := 0
	//foor loop with no arguments simular to while loop
	for {

		fmt.Println("Please enter a number between 1 and 100")
		//.Scan takes user input
		fmt.Scan(&input)

		//test input and output of user and random number.
		//fmt.Println("The number entered is...", input, "the random is", randNum, "the amount of attempts are", i, "Theprevious umber is", sameNumber)

		//conditionals for the user guess
		//too high conditional and increase i by one
		if input > randNum {
			fmt.Println("Youre guess is too high")
			i++
			//too low conditional and i increase by one
		} else if input < randNum {
			fmt.Println("Youre guess is too low")
			i++
			//correct conditional and amount of attempts taken to complete the challenge
		} else {
			fmt.Println("You have guessed the right number it has taken", i, "attempts")
		}
		/*conditional for the same number input from the last attempt
		  the amout of total attemps is decreased by one in this case*/
		if sameNumber == input {
			fmt.Println("you have already picked that number")
			i--
		}
		/*the previous attempt number input is stored in sameNumber
		it only lasts for the duration of one itteration of the loop
		so works for the use of testing against the last attempt*/
		sameNumber = input
		//break out of loop when the number is guessed correctly
		if input == randNum {
			break
		}

	}

}
