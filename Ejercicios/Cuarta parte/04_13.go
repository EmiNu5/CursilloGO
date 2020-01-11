package main

import "fmt"

func main() {
	// 13. Construir una función que dado un string devuelva el mismo string pero sin espacios por delante. (nombrarla leftTrim)

	var unString string = "hola  mundo como estas?  "
	fmt.Print(leftTrim(unString), "final")
}
func leftTrim(unString string) string {
	var i int
	var nuevoString []byte

	for i = 0; i < len(unString); i++ {
		if unString[i] != ' ' {
			break
		}
	}
	for ; i < len(unString); i++ {
		nuevoString = append(nuevoString, unString[i])
	}

	return string(nuevoString)
}
