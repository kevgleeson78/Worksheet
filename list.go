package main

import (
	"container/list"
	"fmt"
)

func minMax() (min int, max int) {
	l := list.New()
	l.PushFront(1)
	l.PushFront(5)
	l.PushFront(13)
	l.PushFront(3)
	l.PushFront(44)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) > max {
			max = e.Value.(int)
		} else if e.Value.(int) < max {
			min = e.Value.(int)
		}
	}

	return min, max
}
func main() {
	min, max := minMax()

	fmt.Println("this is the samllest element in the list", min)
	fmt.Println("this is the largest element in the list", max)
}
