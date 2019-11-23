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
	var i int
	var j int
	var k int
	var monomatriz []int
	var valid bool = true

	for k = 0; k < len(matriz); k++ {
		monomatriz = []int{}
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				monomatriz = append(monomatriz, matriz[i+(k/3*3)][j+(k%3*3)])
			}
		}
		for i = 0; i < len(monomatriz); i++ {
			for j = 1 + i; j < len(monomatriz); j++ {
				if monomatriz[i] == monomatriz[j] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if !valid {
			fmt.Println("repetido 3x3")
		} else {
			fmt.Println("No ta repetido")
		}

		//fmt.Printf("%d \n", monomatriz)
	}

}
