package main

import "fmt"

func main() {
	// 20. Construir una función que dados dos strings devuelva
	//  un entero que represente el índice de la primer ocurrencia del segundo string en el primero.
	//   Si el string no se encuentra devolver -1. (nombrarla findFirstString).
	var primerString string = "programar"
	var segundoString string = "divertido"

	fmt.Println(findFirstString(primerString, segundoString))
}
func findFirstString(primerString string, segundoString string) int {
	var i int
	var j int
	for i = 0; i < len(primerString); i++ {
		for j = 0; j < len(segundoString); j++ {
			if primerString[i] == segundoString[j] {
				return j
				break
			}
		}
	}
	return -1
}
