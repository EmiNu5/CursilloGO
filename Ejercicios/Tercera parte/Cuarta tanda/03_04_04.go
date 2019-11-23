package main

import(
	"fmt"
	"math"
)

func main(){
	// 4. Dada una matriz de 3 x 3 de números enteros, imprimir el menor de cada fila
	var c int //columna
	var f int //fila
	var berta [3][3]int = [3][3]int {
	[3]int {5,8,9},
	[3]int {6,2,3},
	[3]int {1,0,4}}
	var min int

	for f=0;f<3;f++{
		min = math.MaxInt64
		for c=0;c<3;c++{
			if berta[f][c]<min{
				min =berta[f][c]
			}
		}
		fmt.Println(min)
	}
}