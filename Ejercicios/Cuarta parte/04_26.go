package main

import "fmt"

func main26() {
	// 	26. Construir una función que dado un un array de 3 enteros devuelva el mismo array pero
	// 	 invertido (nombrarla reverseInt3).
	// Para pensar: Qué pasaría si en lugar de 3 números enteros, ahora quisiéramos una de 4?
	// y si quisiéramos otra de 5? qué desventaja presenta esto a utilizar un slice?

	var aSlice [3]int = [3]int{0, 4, 5}

	fmt.Printf("%#v", reverseInt3(aSlice))
}
func reverseInt3(aSlice [3]int) [3]int {
	var i int
	var reverse [3]int

	for i = len(aSlice) - 1; i >= 0; i-- {
		reverse[len(aSlice)-i-1] = aSlice[i]
	}
	return reverse
}
