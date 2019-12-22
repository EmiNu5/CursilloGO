package main

import "fmt"

func main() {
	var f int
	var c int
	const filas = 10
	var laberinto [filas][filas]bool = [filas][filas]bool{
		[filas]bool{true, true, true, true, false, false, false, false, false, false},
		[filas]bool{false, false, false, true, false, false, false, false, false, false},
		[filas]bool{false, false, false, true, false, false, false, false, false, false},
		[filas]bool{false, false, false, true, true, true, true, false, false, false},
		[filas]bool{false, false, false, false, false, false, true, false, false, false},
		[filas]bool{false, false, false, false, false, false, true, false, false, false},
		[filas]bool{false, false, false, false, false, false, true, true, true, true},
		[filas]bool{false, false, false, false, false, false, false, false, false, true},
		[filas]bool{false, false, false, false, false, false, false, false, false, true},
		[filas]bool{false, false, false, false, false, false, false, false, false, true},
	}

	for f = 0; f < len(laberinto); f++ {
		fmt.Println("----------------------------------------------------")
		for c = 0; c < len(laberinto); c++ {
			if c == 0 {
				fmt.Print()
				fmt.Print("|")
				if laberinto[f][c] {
					fmt.Print("X    ")
				} else {
					fmt.Print("-    ")
				}
			} else {
				if laberinto[f][c] {
					fmt.Printf("X    ")
				} else {
					fmt.Printf("-    ")
				}
			}
		}
		fmt.Print("|")
		fmt.Println()

	}
	fmt.Println("----------------------------------------------------")
	f = 0
	c = 0
	var fAnt int
	var cAnt int
	var camino []string = []string{"0,0"}
	var fc string
	var salida bool

	for f < len(laberinto) { // LABERINTO
		if c != filas-1 && laberinto[f][c+1] && (cAnt != c+1) {
			/*DERECHA*/
			cAnt = c
			fAnt = f
			c++
			fc = fmt.Sprintf("-> %d,%d", f, c)
			camino = append(camino, fc)
		} else if f != filas-1 && laberinto[f+1][c] && (fAnt != f+1) {
			/*ABAJO*/
			fAnt = f
			cAnt = c
			f++
			fc = fmt.Sprintf("-> %d,%d", f, c)
			camino = append(camino, fc)
		} else if f != 0 && laberinto[f-1][c] && (fAnt != f-1) {
			/*ARRIBA*/
			fAnt = f
			cAnt = c
			f--
			fc = fmt.Sprintf("-> %d,%d", f, c)
			camino = append(camino, fc)
		} else if c != 0 && laberinto[f][c-1] && (cAnt != c-1) {
			/*IZQUIERDA*/
			cAnt = c
			fAnt = f
			c--
			fc = fmt.Sprintf("-> %d,%d", f, c)
			camino = append(camino, fc)
		} else {
			break
		}
	}
	if f == filas-1 && c == filas-1 {
		salida = true
	}
	if !salida {
		fmt.Print("No tiene salida")
	} else {
		fmt.Printf("el camino es: %s", camino)
	}
}
