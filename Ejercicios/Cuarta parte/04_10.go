package main

import "fmt"

func main() {

	//10. Construir una función que dado un string imprima en pantalla el string dado en color rojo. (nombrarla printInRed)

	var unString string = "ESTOY ROJO"

	printInRed(unString)
}
func inRed(unString string) string {
	const red string = "\u001b[31m"
	const reinicio string = "\u001b[0m"

	return red + unString + reinicio
}
func printInRed(unString string) {

	fmt.Printf(inRed(unString))
}
