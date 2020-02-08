package main

import "fmt"

func main25() {
	// 	25. Construir una función que dado un slice de strings y un string devuelva
	//un string con todos los strings del slice concatenados con el string dado. (nombrarla join)
	// Ej:

	// var strings []string = []string{"hola", "que", "tal?"}
	// var joinedString string = join(strings, "--")
	// fmt.Println(joinedString)

	// > "hola--que--tal?"
	var aSlice []string = []string{"hola", "que", "tal"}
	var aString string = "--"

	fmt.Printf("%#v", join(aSlice, aString))
}
func join(aSlice []string, aString string) string {
	var i int
	var palabras string

	for i = 0; i < len(aSlice); i++ {
		if i == 0 {
			palabras = aSlice[i]
		} else {
			palabras = palabras + aString + aSlice[i]
		}
	}
	return palabras
}
