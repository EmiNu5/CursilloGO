package main

import "fmt"

func main() {
	var f int
	var c int
	var matriz [9][9]int = [9][9]int{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[9]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[9]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[9]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
		[9]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
		[9]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
		[9]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
		[9]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
		[9]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
	}
	for f = 0; f < len(matriz); f++ {
		if f%3 == 0 {
			fmt.Println("+ - - - + - - - + - - - +")
		}
		for c = 0; c < len(matriz); c++ {
			if c%3 == 0 {
				fmt.Print()
				fmt.Print("|")
				fmt.Printf("%3d", matriz[f][c])
			} else {
				fmt.Printf("%2d", matriz[f][c])
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("+ - - - + - - - + - - - +")

	// las filas deben contener solo números del 1 al 9
	// las filas no deben contener números repetidos
	// las columnas deben contener solo números del 1 al 9
	// las columnas no debe contener números repetidos
	// cada cuadrante de 3 x 3 debe contener solo números del 1 al 9
	// cada cuadrante de 3 x 3 no debe contener números repetidos

	const minValor int = 1
	const maxValor int = 9
	var numValido bool = true
	var unico bool = true
	var monomatriz []int
	var i int
	var j int

	for f = 0; f < len(matriz); f++ {
		for c = 0; c < len(matriz); c++ {
			if matriz[f][c] < minValor || matriz[f][c] > maxValor {
				numValido = false
				break
			}
		}
		if !numValido {
			break
		}
	}
	if !numValido {
		fmt.Println("Hay valores 1<x>10")
	} else {
		fmt.Println("Valores correctos")
	}
	for f = 0; f < len(matriz); f++ {
		monomatriz = []int{}
		for c = 0; c < len(matriz); c++ {
			monomatriz = append(monomatriz, matriz[f][c])
			if c > 0 || c%9 == 0 {
				for i = 0; i < len(monomatriz); i++ {
					for j = (i + 1); j < len(monomatriz); j++ {
						if monomatriz[i] == monomatriz[j] {
							unico = false
							break
						}
					}
					if !unico {
						break
					}
				}
				if !unico {
					break
				}
			}
			if !unico {
				break
			}
		}
		if !unico {
			fmt.Printf("Filas: numero %d repetido coordenada %d %d .\n", matriz[f][c], f, c)
		} else {
			fmt.Printf("fila %d unica \n", f)
		}
	}
	for f = 0; f < len(matriz); f++ {
		// las columnas no debe contener números repetidos
		monomatriz = []int{}
		for c = 0; c < len(matriz); c++ {
			monomatriz = append(monomatriz, matriz[c][f])
			if c > 0 || c%9 == 0 {
				for i = 0; i < len(monomatriz); i++ {
					for j = (i + 1); j < len(monomatriz); j++ {
						if monomatriz[i] == monomatriz[j] {
							unico = false
							break
						}
					}
					if !unico {
						break
					}
				}
				if !unico {
					break
				}
			}
			if !unico {
				break
			}
		}
		if !unico {
			fmt.Printf(" columnas: numero %d repetido coordenada %d %d .\n", matriz[c][f], c, f)
		} else {
			fmt.Printf("columna %d unica \n", f)
		}
	}
}
