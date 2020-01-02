package main

import "fmt"

func main() {
	var diagonal1 string = "\u2199"   // ↙
	var diagonal2 string = "\u2196"   // ↖
	var horizontal1 string = "\u2190" //←
	var arriba1 string = "\u2191"     //↑
	var player1 string = "e"
	var player2 string = "g"
	var x int
	var y int
	var tablero [5][4]string
	const minValor int = 1
	const maxValor int = 3
	var casillaIncorrecta bool
	var i int
	var ganador bool
	var repetir string
	var gameover bool

	fmt.Println("♦♦♦- BIENVENIDO RANCHO APARTE -♦♦♦ \n")

	for {

		/*DIBUJITO VACIO*/
		for x = 0; x < 3; x++ {
			for y = 0; y < 3; y++ {

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
		for !gameover { /*Entrada y parametros */
			casillaIncorrecta = false
			if i%2 == 0 {

				for !casillaIncorrecta {
					x = 0
					y = 0
					fmt.Printf(" %s ingresa la (X,Y) donde ira la ficha X\n", player1)
					fmt.Scanf("\n%d,%d", &x, &y)
					if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
						fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3\n")
					} else if tablero[x][y-1] == "X" || tablero[x][y-1] == "O" {
						fmt.Println("Ouchh casilla ocupada prueba otra \n")
					} else {
						tablero[x][y-1] = "X"
						casillaIncorrecta = true
						i++
					}
				}
			} else {
				for !casillaIncorrecta {
					x = 0
					y = 0
					fmt.Printf(" %s ingresa la (X,Y) donde ira la ficha O\n", player2)
					fmt.Scanf("\n%d,%d", &x, &y)
					if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
						fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3 \n")
					} else if tablero[x][y-1] == "X" || tablero[x][y-1] == "O" {
						fmt.Println("Ouchh casilla ocupada prueba otra \n ")
					} else {
						tablero[x][y-1] = "O"
						casillaIncorrecta = true
						i++
					}
				}
			}
			/*CASOS DE GANADOR*/
			if tablero[1][0] == tablero[1][1] && tablero[1][0] == tablero[1][2] {
				if tablero[1][0] == "X" || tablero[1][0] == "O" {
					ganador = true
					tablero[1][3] = diagonal1
				}
			} else if tablero[1][0] == tablero[2][0] && tablero[1][0] == tablero[3][0] {
				if tablero[1][0] == "X" || tablero[1][0] == "O" {
					ganador = true
					tablero[4][0] = arriba1
				}

			} else if tablero[1][0] == tablero[2][1] && tablero[1][0] == tablero[3][2] {
				if tablero[1][0] == "X" || tablero[1][0] == "O" {
					ganador = true
					tablero[4][3] = diagonal2
				}
			}

			if tablero[2][1] == tablero[1][1] && tablero[2][1] == tablero[3][1] {
				if tablero[2][1] == "X" || tablero[2][1] == "O" {
					ganador = true
					tablero[4][1] = arriba1
				}
			} else if tablero[2][1] == tablero[1][2] && tablero[2][1] == tablero[3][0] {
				if tablero[2][1] == "X" || tablero[2][1] == "O" {
					ganador = true
					tablero[0][3] = diagonal1
				}
			} else if tablero[2][1] == tablero[2][0] && tablero[2][1] == tablero[2][2] {
				if tablero[2][1] == "X" || tablero[2][1] == "O" {
					ganador = true
					tablero[2][3] = horizontal1
				}
			}

			if tablero[3][2] == tablero[3][1] && tablero[3][2] == tablero[3][0] {
				if tablero[3][2] == "X" || tablero[3][2] == "O" {
					ganador = true
					tablero[3][3] = horizontal1
				}
			} else if tablero[3][2] == tablero[2][2] && tablero[3][2] == tablero[1][2] {
				if tablero[3][2] == "X" || tablero[3][2] == "O" {
					ganador = true
					tablero[4][2] = arriba1
				}

			}

			/*DIBUJITO*/
			for x = 0; x < len(tablero); x++ {
				for y = 0; y < len(tablero[x]); y++ {
					if x == 0 {
						if y < 3 {
							fmt.Printf("%2s  ", tablero[x][y])
						} else {
							fmt.Printf("%2s  \n", tablero[x][y])
						}

					} else if x > 0 && x < 4 {
						if y < 2 {
							fmt.Printf("%2s |", tablero[x][y])
						} else if y > 1 && y < 3 {
							fmt.Printf("%2s  ", tablero[x][y])
						} else if y < 4 {
							fmt.Printf("%2s  \n", tablero[x][y])
						}

					} else if x == 4 {
						if y < 3 {
							fmt.Printf("%2s  ", tablero[x][y])
						} else {
							fmt.Printf("%2s  \n", tablero[x][y])
						}
					}
				}
				if x > 0 && x < 3 {
					fmt.Println("\n-----------")
				}

			}
			if ganador {
				fmt.Print("Has ganado ")
				if i%2 != 0 {
					fmt.Print(player1)
				} else {
					fmt.Print(player2)
				}
			} else if i == 8 {
				fmt.Print("Es un empate")
			}
			for ganador || i == 9 {
				fmt.Println("\n Quieres volver a jugar? s/n ")
				fmt.Scan(&repetir)
				if repetir == "s" || repetir == "n" {
					break
				}
			}

			if repetir == "s" {

				tablero = [5][4]string{}
				gameover = false
				ganador = false
				repetir = ""
				i = 0
				break
			} else if repetir == "n" {
				fmt.Println("Hasta la proxima")
				gameover = true
			}
		}
		if gameover {
			break
		}
	}
}
