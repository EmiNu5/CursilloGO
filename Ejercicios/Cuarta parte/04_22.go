package main

import "fmt"

func main22() {
	// 	22. Construir una función que dado un string y un byte devuelva un slice de strings que contenga todas las particiones del string obtenido de dividirlo por el byte. (nombrarla splitByChar)
	// Ej:

	// var strings []string = splitByChar("Programar es divertido!", ' ')
	// fmt.Println(strings)

	// > ["Programar", "es", "divertido!"]

	var aString string = "Programar es divertido"
	var aByte byte = 'ñ'

	fmt.Printf("%#v", splitByChar(aString, aByte))
}
func splitByChar(aString string, aByte byte) []string {
	var letras []byte
	var palabras []string
	var i int

	for i = 0; i < len(aString); i++ {
		if aString[i] != aByte {
			letras = append(letras, aString[i])
		} else {
			if len(letras) > 0 {
				palabras = append(palabras, string(letras))
				letras = []byte{}
			}
		}
	}
	if len(letras) > 0 {
		palabras = append(palabras, string(letras))
	}

	return palabras
}
