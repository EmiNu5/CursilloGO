package main

import "fmt"

func main() {

	//16. Construir una función que dado un string y un número entero
	//devuelva el string a partir del número usado como índice.
	// Si el índice es inválido devolver un string vacío

	var unString string = "leandrete"
	var num int = 8

	fmt.Printf("%c", indice(unString, num))
}
func indice(unString string, num int) byte {

	if num > len(unString) || num <= 0 {
		return ' '
	}
	return unString[num-1]

}
