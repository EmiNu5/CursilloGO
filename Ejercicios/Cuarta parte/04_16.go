package main

import "fmt"

func main16() {

	//16. Construir una función que dado un string y un número entero
	//devuelva el string a partir del número usado como índice.
	// Si el índice es inválido devolver un string vacío

	var unString string = "leandrete"
	var num int = 6

	fmt.Printf("%s", indice(unString, num))
}
func indice(unString string, num int) string {
	var i int
	var aux []byte

	if num > len(unString) || num <= 0 {
		return ""
	}
	for i = num - 1; i < len(unString); i++ {
		aux = append(aux, unString[i])
	}
	return string(aux)
}
