package main

import "fmt"

func main() {
	// 15. Construir una función que dado un string devuelva el mismo string pero sin espacios por delante ni por detrás. (nombrarla trim)

	var unString string = "   hola  mundo como estas?  "
	fmt.Print(trim(unString), "final")
}
func trim(unString string) string {
	var i int
	var nuevoString []byte

	for i = 0; i < len(unString); i++ {
		if unString[i] != ' ' {
			nuevoString = append(nuevoString, unString[i])
		}

	}
	return string(nuevoString)
}
