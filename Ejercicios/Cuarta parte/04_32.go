package main

import "fmt"

func main32() {
	var bools []bool = []bool{true, false, false, false}

	fmt.Println(anyTrue(bools))
}
func anyTrue(bools []bool) bool {
	var i int
	var rule bool

	for i = 0; i < len(bools); i++ {
		rule = false
		if bools[i] {
			rule = true
			break
		}
	}
	return rule
}
