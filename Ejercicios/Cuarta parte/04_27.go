package main

import "fmt"

func main27() {
	// 	27. Construir una función que dado un slice de números enteros lo devuelva en orden inverso (nombrarla reverseInt).
	var aSlice []int = []int{5, 4, 0}

	fmt.Printf("%#v", reverseInt(aSlice))

}
func reverseInt(aSlice []int) []int {
	var i int
	var reverse []int

	for i = len(aSlice) - 1; i >= 0; i-- {
		reverse = append(reverse, aSlice[i])
	}
	return reverse
}
