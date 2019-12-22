package main

import "fmt"

func main() {
	//Dada una matriz de 3 x 3 de números enteros, imprimir el elemento 2 - 2

	var matriz [3][3]int = [3][3]int{
		[3]int{1, 2, 3},
		[3]int{7, 8, 9},
	}
	fmt.Printf("El elemento 2x2 es: %d", matriz[1][1])
}
