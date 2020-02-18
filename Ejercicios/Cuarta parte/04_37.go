package main

import "fmt"

func main37() {

	var ints []int = []int{33, 0, 5, 4, 1}

	fmt.Println(descOrderInt(ints))
}
func descOrderInt(ints []int) []int {
	var i int
	var j int
	var max int

	for i = 0; i < len(ints); i++ {
		for j = 0; j < len(ints); j++ {
			if ints[i] > ints[j] {
				max = ints[i]
				ints[i] = ints[j]
				ints[j] = max
			}
		}
	}
	return ints
}
