package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var player1 string
	var player2 string
	var x int
	var y int
	var tablero [3][3]string
	const minValor int = 1
	const maxValor int = 3
	var casillaCorrecta bool
	var i int
	var ganador bool
	var repetir string
	var gameover bool

	var horizontal1 string
	var horizontal2 string
	var horizontal3 string
	var vertical1 string
	var vertical2 string
	var vertical3 string
	var diagonalarriba string
	var diagonalabajo string

	var jugadorActual string
	var fichaActual string

	var pc string
	rand.Seed(time.Now().Unix())

	fmt.Println("♦♦♦- BIENVENIDO -♦♦♦ \n")
	for {
		fmt.Println("Desea jugar contra la PC?? s/n")
		fmt.Scan(&pc)
		if pc == "s" || pc == "n" {
			break
		}
	}
	if pc == "n" {
		fmt.Println("Player1 Ingresa tu nombre: ")
		fmt.Scan(&player1)
		fmt.Println("Player2 ingresa tu nombre: ")
		fmt.Scan(&player2)
	} else {
		fmt.Println("Player1 Ingresa tu nombre: ")
		fmt.Scan(&player1)
		player2 = "PC"
	}

	fmt.Printf("%s vs %s \n\n", player1, player2)

	for {

		/*DIBUJITO VACIO*/
		for x = 0; x < len(tablero); x++ {
			for y = 0; y < len(tablero); y++ {

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

			if i%2 == 0 {
				jugadorActual = player1
				fichaActual = "X"
			} else {
				jugadorActual = player2
				fichaActual = "O"
			}

			casillaCorrecta = false
			for !casillaCorrecta {
				fmt.Printf(" %s ingresa la (X,Y) donde ira la ficha X\n", jugadorActual)
				if i%2 != 0 && pc == "s" {
					for {
						x = rand.Intn(maxValor-minValor) + 1
						y = rand.Intn(maxValor-minValor) + 1
						if tablero[x-1][y-1] != "X" || tablero[x-1][y-1] != "O" {
							break
						}
					}
					time.Sleep(2 * time.Second)
				} else {
					fmt.Scanf("\n%d,%d", &x, &y)
				}
				if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
					fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3\n")
				} else if tablero[x-1][y-1] == "X" || tablero[x-1][y-1] == "O" {
					fmt.Println("Ouchh casilla ocupada prueba otra \n")
				} else {
					tablero[x-1][y-1] = fichaActual
					casillaCorrecta = true
				}
			}

			/*CASOS DE GANADOR*/
			if tablero[0][0] == tablero[0][1] && tablero[0][0] == tablero[0][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
					horizontal1 = "\u2190"

				}
			} else if tablero[0][0] == tablero[1][0] && tablero[0][0] == tablero[2][0] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
					vertical1 = "\u2191"
				}

			} else if tablero[0][0] == tablero[1][1] && tablero[0][0] == tablero[2][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
					diagonalarriba = "\u2196"
				}
			}

			if tablero[1][1] == tablero[0][1] && tablero[1][1] == tablero[2][1] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
					vertical2 = "\u2191"
				}
			} else if tablero[1][1] == tablero[0][2] && tablero[1][1] == tablero[2][0] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
					diagonalabajo = "\u2199"
				}
			} else if tablero[1][1] == tablero[1][0] && tablero[1][1] == tablero[1][2] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
					horizontal2 = "\u2190"
				}
			}

			if tablero[2][2] == tablero[2][1] && tablero[2][2] == tablero[2][0] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
					horizontal3 = "\u2190"
				}
			} else if tablero[2][2] == tablero[1][2] && tablero[2][2] == tablero[0][2] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
					vertical3 = "\u2191"
				}

			}

			/*DIBUJITO*/
			fmt.Printf("             %1s\n", diagonalabajo)
			fmt.Printf(" %1s | %1s | %1s %1s\n", tablero[0][0], tablero[0][1], tablero[0][2], horizontal1)
			fmt.Println("-----------")
			fmt.Printf(" %1s | %1s | %1s %1s\n", tablero[1][0], tablero[1][1], tablero[1][2], horizontal2)
			fmt.Println("-----------")
			fmt.Printf(" %1s | %1s | %1s %1s\n", tablero[2][0], tablero[2][1], tablero[2][2], horizontal3)
			fmt.Printf(" %1s   %1s   %1s %1s\n", vertical1, vertical2, vertical3, diagonalarriba)

			if ganador {
				fmt.Print("Has ganado ")
				if i%2 == 0 {
					fmt.Print(player1)
				} else {
					fmt.Print(player2)
				}
			} else if i == 9 {
				fmt.Print("Es un empate")
			}
			for ganador || i == 9 {
				fmt.Println("\n Quieres volver a jugar? s/n ")
				fmt.Scan(&repetir)
				if repetir == "s" || repetir == "n" {
					break
				}
			}
			i++

			if repetir == "s" {

				tablero = [3][3]string{}
				gameover = false
				ganador = false
				repetir = ""
				horizontal1 = ""
				horizontal2 = ""
				horizontal3 = ""
				vertical1 = ""
				vertical2 = ""
				vertical3 = ""
				diagonalarriba = ""
				diagonalabajo = ""
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
