package main

import "fmt"

//21. Construir una función que dados dos strings
// devuelva un slice de enteros que contengan los índices de todas las ocurrencias
// del segundo string en el primero (nombrarla findAllString).

func main21() {
	var firstString string = "papapas"
	var secondString string = "papapa"

	fmt.Println(findAllString(firstString, secondString))
}
func findAllString(firstString string, secondString string) []int {
	var i int
	var j int
	var found bool
	var ocurrence []int = []int{}

	for i = 0; i < len(firstString); i++ {
		found = true
		for j = 0; j < len(secondString); j++ {
			if j+i >= len(firstString) || firstString[i+j] != secondString[j] {
				found = false
				break
			}
		}
		if found {
			ocurrence = append(ocurrence, i)
		}
	}
	return ocurrence
}
