package main

import "fmt"

// // 2. Dada una matriz de 3 x 3 de números enteros, imprimir toda la matriz de la siguiente forma
//      c1 c2 c3  
//     ---------- 
// f1 | n1 n2 n3 |
// f2 | n4 n5 n6 |
// f3 | n7 n8 n9 |
// ---------- 
func main(){
	var f int //fila
	var c int //columna
	var berta [3][3]int = [3][3]int {
	[3]int {5,8,9},
	[3]int {6,2,3},
	[3]int {1,0,4}}

	fmt.Println("   c1 c2 c3")
	fmt.Println("   --------")
	for f = 0;f<3;f++{
		for c=-1;c<3;c++{
			if  c == -1 {
				fmt.Printf("f%d|",f+1)
			} else{
				fmt.Printf("%2d",berta[f][c])
			}
		}
		fmt.Println(" |")
	}
	fmt.Println("   --------")
}