package main

import "fmt"

func main() {
	// 19. Construir una función que dado un string y un byte devuelva
	//  un slice de enteros que represente todos los índices con las ocurrencias de ese byte en el string
	//  (nombrarla findAllChar).
	var unString string = "jijiji"
	var unChar byte = 'i'

	fmt.Printf("%d", findAllChar(unString, unChar))
}
func findAllChar(unString string, unChar byte) []int {
	var unSlice []int
	var i int

	for i = 0; i < len(unString); i++ {
		if unString[i] == unChar {
			unSlice = append(unSlice, i)
		}
	}

	return unSlice
}
