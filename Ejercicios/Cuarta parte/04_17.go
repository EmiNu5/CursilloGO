package main

import "fmt"

func main() {
	// 17. Contruir una función que dado un string y dos números enteros
	// devuelva el string a partir del primer número usado como índice de comienzo y
	//  usando el segundo como largo del string. Si el largo es mayor al posible restante, devolver el string hasta el final.
	//  Si el índice es inválido devolver un string vacío. (nombrarla substr).

	var unString string = "holalea"
	var inicio int = 3
	var largo int = 80

	fmt.Print(substr(unString, inicio, largo))
}
func substr(unString string, inicio int, largo int) string {

	var nuevoString []byte
	var i int

	if inicio <= 0 || inicio >= len(unString) {
		return "un string vacío"
	}
	if largo+inicio > len(unString) {
		for i = inicio - 1; i < len(unString)-1; i++ {
			nuevoString = append(nuevoString, unString[i])
		}
	} else {
		for i = inicio - 1; i < inicio+largo-1; i++ {
			nuevoString = append(nuevoString, unString[i])
		}
	}
	return string(nuevoString)
}
