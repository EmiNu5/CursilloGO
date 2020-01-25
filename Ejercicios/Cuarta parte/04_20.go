package main

import "fmt"

func main() {
	// 20. Construir una función que dados dos strings devuelva
	//  un entero que represente el índice de la primer ocurrencia del segundo string en el primero.
	//   Si el string no se encuentra devolver -1. (nombrarla findFirstString).
	var primerString string = "programar"
	var segundoString string = "mar"

	fmt.Println(findFirstString(primerString, segundoString))
}
func findFirstString(primerString string, segundoString string) int {
	var i int
	var j int
	var coincidencia int = 0
	var indice int = -1

	for i = 0; i < len(primerString); i++ {
		fmt.Println("indice", i)
		for ; j < len(segundoString); j++ {
			if primerString[i] == segundoString[j] {
				coincidencia++
			} else {
				break
			}
		}
		fmt.Println("coincidencia", coincidencia)
		if coincidencia == len(segundoString) {
			fmt.Println(coincidencia, "==", len(segundoString))
			indice = len(primerString) - coincidencia
			break
		}

	}
	return indice
}
