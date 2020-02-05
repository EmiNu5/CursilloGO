package main

import "fmt"

func main15() {
	// 15. Construir una función que dado un string devuelva el mismo string
	// pero sin espacios por delante ni por detrás. (nombrarla trim)

	var str1 string = "  hola como andas?   "

	fmt.Println(leftTrim(rightTrim(str1)) + ".")
}
