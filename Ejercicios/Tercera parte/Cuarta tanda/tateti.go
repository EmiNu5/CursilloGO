package main

import "fmt"

func main() {

	var player1 string
	var player2 string
	var x int
	var y int
	var tablero [3][3]string
	const minValor int = 1
	const maxValor int = 3
	var limite bool
	var ocupada bool
	var i int

	fmt.Println("♦♦♦- BIENVENIDO RANCHO APARTE -♦♦♦ \n")
	fmt.Println("Player1 Ingresa tu nombre: ")
	fmt.Scan(&player1)
	fmt.Println("Player2 ingresa tu nombre: ")
	fmt.Scan(&player2)
	fmt.Printf("%s vs %s \n\n", player1, player2)

	/*DIBUJITO VACIO*/
	for x = 0; x < len(tablero); x++ {
		for y = 0; y < len(tablero); y++ {
			tablero[x][y] = ""
			if y < 2 {
				fmt.Print("   |")
			} else {
				fmt.Println("   ")
			}
		}
		if x < 2 {
			fmt.Println("-----------")
		}
	}
	for { /*FOR INFINITO*/
		limite = false
		ocupada = false
		if i%2 == 0 {
			for !limite && !ocupada {
				fmt.Printf(" %s Ingrese la (X,Y) donde ira la ficha X\n", player1)
				fmt.Scan(&x, &y)
				if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
					fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3")
				} else if tablero[x-1][y-1] == "X" || tablero[x-1][y-1] == "O" {
					fmt.Println("Ouchh casilla ocupada prueba otra ;) ")
				} else {
					tablero[x-1][y-1] = "X"
					limite = true
					ocupada = true
					i++
				}
			}
		} else {
			for !limite && !ocupada {
				fmt.Printf(" %s Ingrese la (X,Y) donde ira la ficha O\n", player2)
				fmt.Scan(&x, &y)
				if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
					fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3")
				} else if tablero[x-1][y-1] == "X" || tablero[x-1][y-1] == "O" {
					fmt.Println("Ouchh casilla ocupada prueba otra ;) ")
				} else {
					tablero[x-1][y-1] = "O"
					limite = true
					ocupada = true
					i++
				}
			}
		}
		/*DIBUJITO*/
		for x = 0; x < len(tablero); x++ {
			for y = 0; y < len(tablero); y++ {
				if y < 2 {
					fmt.Printf("%2s |", tablero[x][y])
				} else {
					fmt.Printf(" %2s \n", tablero[x][y])
				}
			}
			if x < 2 {
				fmt.Println("-----------")
			}
		}

	}
}
