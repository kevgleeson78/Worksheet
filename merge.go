//worksheet problem 8
//Merge and sort two arrays
//@Author Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet
//date 19/09/2017
//source https://golang.org/pkg/sort/
package main

//import fmt and sort packages
import (
	"fmt"
	"sort"
)

//A function that takes in two unordered lists
//then returns a merged and ordered list from the two lists.
func merge(list1 []int, list2 []int) (list3 []int) {
	
	//Sort method to sort each list
	sort.Ints(list1)
	sort.Ints(list2)
	//append the two lists together and store in varible list3
	list3 = append(list1, list2...)
	//Sort list3
	sort.Ints(list3)
	//return list3
	return  list3
}
func main() {
	//store the three lists for printing out to the screen
	//Two arrays of different sizes and unordered
	list1 := []int{22, 18, 1, 78, 99, 4}
	list2 := []int{11, 32, 41, 5, 204, 71, 2}
	
	//Print out the result of the function
	fmt.Println("The first sorted list is ", list1, "\nThe Second sorted list is ", list2, "\nThe mreged and ordered list is ", merge(list1, list2))

}
