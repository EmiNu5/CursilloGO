package main

import "fmt"

func main24() {
	// 24. Construir una función que dados dos strings devuelva un bool dependiendo si el primer string contiene al segundo.
	//  (nombrarla contains)

	var firstString string = "mateico"
	var secondString string = "mate"

	fmt.Printf("%#v", contains2(firstString, secondString))

}
func contains(firstString string, secondString string) bool {
	var i int
	var j int
	var found bool

	if len(secondString) != 0 {
		for i = 0; i < len(firstString); i++ {
			found = true
			for j = 0; j < len(secondString); j++ {
				if j+i >= len(firstString) || firstString[i+j] != secondString[j] {
					found = false
					break
				}
			}
			if found {
				break
			}
		}
	} else {
		found = false
	}

	return found
}

func contains2(aString string, searchedString string) bool {
	var contais bool = findFirstString(aString, searchedString) != -1
	return contais
}
