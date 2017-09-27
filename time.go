//worksheet problem 2
//@Author KEvin Gleeson
//Scource https://tour.golang.org/welcome/4

package main

import (
	"fmt"
	//import the time.Now() method
	"time"
)

//curley brace needs to be directly beside the functions brackets
func main() {
	fmt.Println("this function will display the currrent time")

	fmt.Println("The currrent time is", time.Now())
}
