package main

import "fmt"

func main() {
	// 15. Construir una función que dado un string devuelva el mismo string
	// pero sin espacios por delante ni por detrás. (nombrarla trim)

	var str1 string = "  hola como andas?   "

	fmt.Println(leftTrim(rightTrim(str1)) + ".")
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
