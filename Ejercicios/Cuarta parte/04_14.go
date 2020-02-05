package main

import "fmt"

func main14() {
	// 14. Construir una función que dado un string devuelva el mismo string pero sin espacios por detrás. (nombrarla rightTrim)

	var unString string = "hola  mundo como estas?  "
	fmt.Print(rightTrim(unString), "final")
}
func rightTrim(unString string) string {
	var i int
	var nuevoString []byte
	var j int

	for i = len(unString) - 1; i >= 0; i-- {
		if unString[i] != ' ' {
			break
		}
	}
	for j = 0; j <= i; j++ {
		nuevoString = append(nuevoString, unString[j])
	}

	return string(nuevoString)
}
