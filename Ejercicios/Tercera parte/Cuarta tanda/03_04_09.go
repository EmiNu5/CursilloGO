package main

import "fmt"

func main() {
	const filas int = 10
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
	var f int
	var c int
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
