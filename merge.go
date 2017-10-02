//worksheet problem 8
//@Author Kevin Gleeson
//date 19/09/2017

package main

//import fmt and sort packages
import (
	"fmt"
	"sort"
)

//A function that merges two lists and puts them in order
func merge() (list1 []int, list2 []int, list3 []int) {
	//Two arrays of different sizes and unordered
	list1 = []int{22, 18, 1, 78, 99, 4}
	list2 = []int{11, 32, 41, 5, 204, 71, 2}
	//Sort method to sort each list
	sort.Ints(list1)
	sort.Ints(list2)
	//append the two lists together and store in varible list3
	list3 = append(list1, list2...)
	//Sort list3
	sort.Ints(list3)
	//return list3
	return list2, list1, list3
}
func main() {
	//store the three lists for printing out to the screen
	list1, list2, list3 := merge()
	//Print out the result of the function
	fmt.Println("The first sorted list is ", list1, "\nThe Second sorted list is ", list2, "\nThe mreged and ordered list is ", list3)

}
