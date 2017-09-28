package main

import (
	"fmt"
	"sort"
)

func merge(list1 []int, list2 []int) (int, error) {
	sort.Ints(list1)
	sort.Ints(list2)
	list3 := append(list1, list2...)

	return fmt.Println(list3)
}
func main() {
	list1 := []int{22, 18, 1, 78, 99, 4}
	list2 := []int{11, 32, 41, 5, 204, 71, 2}
	merge(list1, list2)

}
