package main

import "fmt"

func main(){
	//Dada una matriz de 3 x 3 de números enteros, imprimir el elemento 2 - 2
	const filas int = 3
	const columnas int = 3
	const largo int = 12
	var i int

	var matrix [filas][columnas]string = [filas][columnas]string {[filas]string{"n1","n2","n3"},[filas]string{"n4","n5","n6"},[filas]string{"n7","n8","n9"}}

	// 2. Dada una matriz de 3 x 3 de números enteros, imprimir toda la matriz de la siguiente forma
	// 		c1 c2 c3  
	// 	   ---------- 
	// f1 | n1 n2 n3 |
	// f2 | n4 n5 n6 |
	// f3 | n7 n8 n9 |
	//     ---------- 

	for i=0;i<largo;i++{
		fmt.Printf("-")
		
	}
	fmt.Println()

	for i=0;i<len(matrix);i++{
		for j=0;j<len(matrix);j++{
			if (matrix + j)/2 
		}
		fmt.Println(matrix[i])
		
	}

	for i=0;i<largo;i++{
		fmt.Printf("-")
		
	}

}