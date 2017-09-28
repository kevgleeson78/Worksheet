package main

import (
	"fmt"
	"sort"
)

func merge() (list3 []int) {
	list1 := []int{22, 18, 1, 78, 99, 4}
	list2 := []int{11, 32, 41, 5, 204, 71, 2}
	sort.Ints(list1)
	sort.Ints(list2)
	list3 = append(list1, list2...)

	return list3
}
func main() {

	result := merge()

	fmt.Println(result)

}
