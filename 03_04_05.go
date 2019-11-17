package main

import (
	"fmt"
	"math"
)
// 5. Dada una matriz de 3 x 3 de número enteros, imprimir el menor de cada columna
func main(){

	var f int // fila
	var c int //columna
	var berta [3][3]int = [3][3]int {
	[3]int {5,8,9},
	[3]int {6,2,3},
	[3]int {1,0,4}}
	var min int

	for f=0;f<3;f++{
		min = math.MaxInt64
		for c=0;c<3;c++{
			if berta[c][f]<min{
				min =berta[c][f]
			}
		}
		fmt.Println(min)
	}
}