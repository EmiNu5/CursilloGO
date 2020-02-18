package main

import (
	"fmt"
)

func main36() {
	var ints []int = []int{33, 0, 5, 4, 1}

	fmt.Println(ascOrderInt(ints))
}
func ascOrderInt(ints []int) []int {
	var i int
	var j int
	var min int

	for i = 0; i < len(ints); i++ {
		for j = 0; j < len(ints); j++ {
			if ints[i] < ints[j] {
				min = ints[i]
				ints[i] = ints[j]
				ints[j] = min
			}
		}
	}
	return ints
}
