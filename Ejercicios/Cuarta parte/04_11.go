package main

import "fmt"

func main() {

	//11. Construir una función que dado un string imprima en pantalla el string dado en color amarillo. (nombrarla printInYellow)

	var unString string = "ESTOY AMARILLO"

	printInYellow(unString)
}
func inYellow(unString string) string {
	const yellow string = " \u001b[33m"
	const reinicio string = "\u001b[0m"

	return yellow + unString + reinicio
}
func printInYellow(unString string) {

	fmt.Printf(inYellow(unString))
}
