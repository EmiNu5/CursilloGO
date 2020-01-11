package main

import "fmt"

func main() {

	var numero1 int = 78
	var numero2 int = 4949

	fmt.Print(sumar(numero1, numero2))
}
func sumar(numero1 int, numero2 int) int {
	var total int

	total = numero1 + numero2

	return total
}
