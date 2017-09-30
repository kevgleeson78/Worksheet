//Problem sheet 9
//Newtons Method
//@author Kevin Gleeson 30/09/17

package main

//import math and fmt
import "fmt"
import "math"

//recursive funtion to take in inital values for x and z

func newton(x, z float64) float64 {
	// formula fro newtons method
	z_next := z - ((z*z - x) / (2 * z))
	//print out each result of the calculation
	fmt.Printf("|%0v|\n", z_next)
	/*conditional to check if the current number
	is equal to the square root of the original number*/
	if z_next == math.Sqrt(x) {
		//return value if it is equal
		return z_next
	}
	/*replace the result of z with z_next
	back into the function
	this repeats until the condition is met*/
	return newton(x, z_next)
}

func main() {
	//initialis z to 1.0
	z := 1.0
	//send first entries to function
	newton(3333333, z)

}
