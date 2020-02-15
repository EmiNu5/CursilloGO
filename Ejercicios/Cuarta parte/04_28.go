package main

import "fmt"

func main28() {
	var aSlice []string = []string{"a", "m", "a", "r"}

	fmt.Printf("%#v", reverseString(aSlice))
}
func reverseString(aSlice []string) []string {
	var i int
	var reverse []string

	for i = len(aSlice) - 1; i >= 0; i-- {
		reverse = append(reverse, aSlice[i])
	}
	return reverse
}
