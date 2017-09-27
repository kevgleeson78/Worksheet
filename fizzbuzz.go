//worksheet problem 3
//@Author Kevin Gleeson
//date 19/09/2017

package main

import "fmt"

func main() {

	// for loop to initialise i to zero and count to 100
	for i := 1; i <= 100; i++ {
		//store mod3 and mod5 in variables
		mod3 := i%3 == 0
		mod5 := i%5 == 0
		//condition for multiples of 3 and 5
		if mod3 && mod5 {
			fmt.Println("fizzbuz")
			//condition for multiples of three
		} else if mod3 {
			fmt.Println("Fizz")
			//condition for multiples of 5
		} else if mod5 {
			fmt.Println("Buzz")
			//condition to print out the remianing numbers
		} else {
			fmt.Println(i)
		}

	}

}
