package main

import "fmt"

func main() {
	// 20. Construir una función que dados dos strings devuelva
	//  un entero que represente el índice de la primer ocurrencia del segundo string en el primero.
	//   Si el string no se encuentra devolver -1. (nombrarla findFirstString).
	var primerString string = "murcielago"
	var segundoString string = "lago"

	fmt.Println(findFirstString(primerString, segundoString))
}
func findFirstString(primerString string, segundoString string) int {
	var i int
	var j int
	var coincidencia int = 0
	var indice int = -1

	if len(segundoString) <= len(primerString) {
		for i = 0; i < len(primerString); i++ {
			fmt.Println("i = ", i)
			for j = 0; j < len(segundoString); j++ {
				if primerString[i+j] == segundoString[j] {
					coincidencia++
				} else {
					break
				}
			}
			fmt.Println("encontradas", coincidencia)
			if coincidencia == len(segundoString) {
				fmt.Println(coincidencia, "==", len(segundoString))
				indice = i
				break
			}

		}
	}
	return indice
}
