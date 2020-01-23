package main

import "fmt"

func main() {

	// 18. Construir una función que dado un string y un byte devuelva
	//  un entero que represente el índice de la primer ocurrencia de ese byte en el string.
	//   Si el byte no se encuentra devolver -1. (nombrarla findFirstChar).

	var unString string = "programar es divertido"
	var unChar byte = 'd'

	fmt.Println(findFirstChar(unString, unChar))
}
func findFirstChar(unString string, unChar byte) int {
	var i int
	var indiceEncontrado int = -1

	for i = 0; i < len(unString); i++ {
		if unString[i] == unChar {
			indiceEncontrado = i - 1
			break
		}
	}
	return indiceEncontrado
}
