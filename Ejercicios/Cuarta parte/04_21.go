package main

import "fmt"

func main() {
	//21. Construir una función que dados dos strings devuelva
	// un slice de enteros que contengan los índices de todas las ocurrencias del segundo string en el primero
	// (nombrarla findAllString).

	var str1 string = "papanatas"
	var str2 string = "pa"

	fmt.Printf(findAllString(str1, str2))

}
func findAllString(str1 string, str2 string) []int {
	var unSlice []int
	var i int
	var j int
	var aux int

	for i = 0; i < len(str1); i++ {
		for j = i; j < len(str2); j++ {
			if str2[j] == str1[i] {
				i++
			}
		}
	}
}
