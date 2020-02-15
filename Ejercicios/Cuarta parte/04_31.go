package main

import "fmt"

func main31() {
	var bools []bool = []bool{true, false, true, true}

	fmt.Println(allTrue(bools))
}
func allTrue(bools []bool) bool {
	var i int
	var rule bool = true

	for i = 0; i < len(bools); i++ {
		if !bools[i] {
			rule = false
			break
		}
	}
	return rule
}
