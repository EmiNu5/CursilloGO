package main

import "fmt"

func main23() {
	// 	23. Construir una función que dados dos strings devuelva un slice de strings que contenga todas las particiones del primer string obtenido de dividirlo por el segundo string. (nombrarla splitByString)
	// Ej:

	// var strings []string = splitByString("la cama la puerta y la mesa", "la")
	// fmt.Println(strings)

	// > [" cama ", " puerta y ", " mesa"]
	var firstString string = "emiliano"
	var secondString string = "ano"

	fmt.Printf("%#v", splitByString(firstString, secondString))
}
func splitByString(firstString string, secondString string) []string {
	var palabras []string
	var letras []byte
	var found bool
	var i int

	for i = 0; i < len(firstString); i++ {
		found = exist(firstString, secondString, i)
		if found {
			if len(letras) > 0 {
				palabras = append(palabras, string(letras))
				letras = []byte{}
			}
			i = i + len(secondString) - 1
		} else {
			letras = append(letras, firstString[i])
		}
	}
	if len(letras) > 0 {
		palabras = append(palabras, string(letras))
	}
	return palabras
}
func exist(firstString string, secondString string, i int) bool {
	var equal bool
	var j int

	equal = true
	if len(secondString) != 0 {
		for j = 0; j < len(secondString); j++ {
			if i+j >= len(firstString) || firstString[i+j] != secondString[j] {
				equal = false
				break
			}

		}
	} else {
		equal = false
	}
	return equal

}
