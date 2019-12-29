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
	var ganador bool
	var repetir string = "s"

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
	for {
		if repetir == "s" {

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

			limite = false
			ocupada = false
			if repetir == "s" && i%2 == 0 {
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
			if tablero[0][0] == tablero[0][1] && tablero[0][0] == tablero[0][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}
			} else if tablero[0][0] == tablero[1][0] && tablero[0][0] == tablero[2][0] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}

			} else if tablero[0][0] == tablero[1][1] && tablero[0][0] == tablero[2][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}
			}
			if tablero[1][1] == tablero[0][1] && tablero[1][1] == tablero[2][1] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			} else if tablero[1][1] == tablero[0][2] && tablero[1][1] == tablero[2][0] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			} else if tablero[1][1] == tablero[1][0] && tablero[1][1] == tablero[1][2] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			}

			if tablero[2][2] == tablero[2][1] && tablero[2][2] == tablero[2][0] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
				}
			} else if tablero[2][2] == tablero[1][2] && tablero[2][2] == tablero[0][2] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
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
			if ganador {
				fmt.Print("Has ganado ")
				if i%2 != 0 {
					fmt.Printf("%s ", player1)
					for {
						fmt.Println("Repetimos?? \n-- ingresa s para si y n para salir")
						fmt.Scan(&repetir)
						if repetir == "s" || repetir == "n" {
							break
						}
					}

				} else {
					fmt.Printf("%s ", player2)
					for {
						fmt.Println("Repetimos?? \n-- ingresa s para si y n para salir")
						fmt.Scan(&repetir)
						if repetir == "s" || repetir == "n" {
							break
						}
					}
				}
			} else if i == 9 {
				fmt.Println("Empateeee weeee ")
				for {
					fmt.Println("Repetimos?? \n Ingresa s para si y n para salir")
					fmt.Scan(&repetir)
					if repetir == "s" || repetir == "n" {
						break
					}
				}

			}
			if repetir == "n" {
				fmt.Println("Adios")
				break
			}
		}
	}
}
