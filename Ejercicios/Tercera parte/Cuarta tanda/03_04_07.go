package main

import (
	"fmt"
)

// 7. Dada una matriz de 3 x 3 de número enteros, imprimir si representa un cuadrante válido para el SUDOKU según las siguientes reglas
// los números de cada casillero deben estar entre el 1 y el 9
// no deben repetirse números en ningún casillero de la matriz

func main(){
	
	var f int
	var c int
	var matriz [3][3]int = [3][3]int{
	[3]int{1,8,3}, 
	[3]int{4,6,6},
	[3]int{7,2,9},
	}
	var monomatriz []int
	const minValor int = 1
	const maxValor int = 10
	var unico bool = true
	var numValido bool = true

	for f=0;f<len(matriz);f++{
		for c=0;c<len(matriz[f]);c++{
			monomatriz = append(monomatriz,matriz[f][c])
			if matriz[f][c] < minValor || matriz[f][c] > maxValor{
				numValido = false
				break
			} 
		}
		if !numValido{
			break
		}
	}
	for f=0;f<len(monomatriz);f++{
		for c=f+1;c<len(monomatriz);c++{
			if monomatriz[f] == monomatriz[c]{
				unico = false
				break
			}
		}
		if !unico{
			break
		}
	}
	if !numValido || !unico {
		fmt.Println("No cumple las reglas")
	} else{
		fmt.Println("Representa un cuadrante unico")
	}

}