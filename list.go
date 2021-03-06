//worksheet problem 6
//Largest and smalest elements in list
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet
//date 19/09/2017
//source https://golang.org/pkg/container/list/
package main

import (
	"container/list"
	"fmt"
)

//function to take in two perameters
func minMax() (min int, max int) {
	//new list created
	l := list.New()
	//populate list with numbers
	l.PushFront(1)
	l.PushFront(5)
	l.PushFront(13)
	l.PushFront(3)
	l.PushFront(44)
	//for loop to get the largest and smallest number in the list
	for e := l.Front(); e != nil; e = e.Next() {
		//conditionals to check each number and if it is higher than the last number checked.

		//e.Value to get the numeric value of the list elements.
		if e.Value.(int) > max {
			max = e.Value.(int)
		} else if e.Value.(int) < max {
			min = e.Value.(int)
		}
	}
	//return result
	return min, max
}
func main() {
	//store results from function
	min, max := minMax()
	//print out the results

	fmt.Println("this is the samllest element in the list", min)
	fmt.Println("this is the largest element in the list", max)
}
