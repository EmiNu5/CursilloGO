package main

import "fmt"

func main() {

	//9. Construir una función que dado un string imprima en pantalla el string dado en color verde. (nombrarla printInGreen)

	var unString string = "ESTOY VERDE"

	printInGreen(unString)
}
func inGreen(unString string) string {
	const green string = "\u001b[32m"
	const reinicio string = "\u001b[0m"

	return green + unString + reinicio
}
func printInGreen(unString string) {

	fmt.Printf(inGreen(unString))
}
