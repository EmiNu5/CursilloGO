package main

import "fmt"

func main10() {

	//10. Construir una función que dado un string imprima en pantalla el string dado en color rojo. (nombrarla printInRed)

	var unString string = "ESTOY ROJO"

	printInRed(unString)
}
func printInRed(unString string) {

	fmt.Printf(inRed(unString))
}
