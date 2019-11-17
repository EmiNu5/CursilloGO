package main

import (
	"fmt"
	"math"
)

// Dada una matriz de 3 x 3 de números enteros, imprimir el mayor de toda la matriz

func main(){
	var f int //fila
	var c int //columna
	var berta [3][3]int = [3][3]int {
	[3]int {5,8,9},
	[3]int {6,2,3},
	[3]int {1,0,4}}

	var max int = math.MinInt64

	for f=0;f<3;f++{
		for c=0;c<3;c++{
			if berta[f][c] > max{
				max = berta[f][c]
			}
		}
	}
	fmt.Println(max)
}