package main

import "fmt"

func main() {
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
	var f int
	var c int
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

	var contadorFila [9]int
	//var contadorColumna [9]int
	var valid bool

	for f = 0; f < 9; f++ {
		for c = 0; c < 9; c++ {
			// contadorColumna[(matriz[c][f])-1]++
			// if contadorColumna[(matriz[c][f])-1] > 1 {
			// 	valid = false
			// 	break
			// }
			contadorFila[(matriz[f][c])-1]++
			if contadorFila[(matriz[f][c])-1] > 1 {
				valid = false
				break
			}
		}
	}
	if !valid {
		fmt.Print("no cumple las reglas")
	} else {
		fmt.Print("cumple")
	}

}
