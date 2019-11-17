package main

import (
	"fmt"
	
)
// 6. Dada una matriz de 3 x 3 de número enteros, imprimir si representa un cuadrante válido según las siguientes reglas
// la sumatoria de toda la matriz no debe exceder el valor 50
// la sumatoria de cada fila no debe exceder el valor 30
// la sumatoria de cada columna no debe ser menor al valor 10
// ningún valor de la matriz debe exceder el valor 10
// ningún valor de la matriz debe ser menor al valor 1

func main(){
	
	var f int // fila
	var c int //columna
	var berta [3][3]int = [3][3]int {
	[3]int {3,5,5},
	[3]int {1,4,2},
	[3]int {6,1,3}}
	var acumMatriz int
	const limitAcumMatriz int=50
	var acumFila int
	const limitAcumFila int=30
	var maxValue int
	const limitMaxValue int=10
	var minValue int
	const limitMinValue int=0
	var acumColum int
	const limitAcumColum int = 10
	var valid bool
		
	for f=0;f<len(berta);f++{
		valid = true
		acumFila=0
		acumColum=0	
		for c=0;c<len(berta[f]);c++{

			acumMatriz= (acumMatriz + berta[f][c])

			acumFila= (acumFila + berta[f][c])
			if acumFila > limitAcumFila {
				valid = false
				break
			}
			if berta[f][c] > limitMaxValue{
				valid = false
				maxValue = berta[f][c]
				break
			}
			if berta[f][c] < limitMinValue{
				valid = false
				minValue = berta[f][c]
				break
			}
			acumColum = (acumColum + (berta[c][f]))
			
		}
		if !valid {
			break
		}
		if acumColum < limitAcumColum {
			valid = false
			break
		} 

	}
	if acumMatriz>limitAcumMatriz{
		fmt.Println(acumMatriz,"La sumatoria de matriz supera el valor 50.")
	} else if acumFila >limitAcumFila {
		fmt.Println(acumFila,"La sumatoria de fila supera el valor 30.")
	} else if maxValue > limitMaxValue{
		fmt.Println(maxValue,"Ningun numero puede exeder el valor 10.")
	} else if minValue < limitMinValue{
		fmt.Println(minValue,"Ningun numero puede ser menor a 1.")
	} else if acumColum < limitAcumColum{
		fmt.Println(acumColum,"La sumatoria de columna no puede ser menor a 10.")
	} else{
		fmt.Println("Cumple las reglas")
	}
}