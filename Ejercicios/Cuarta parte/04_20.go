package main

import "fmt"

func main20() {
	fmt.Println(findFirstString("El primer string va a ser 17", "va"))
}

func findFirstString(aString string, searchedString string) int {
	var found bool
	var i int
	var j int

	for i = 0; i < len(aString); i++ {
		found = true
		for j = 0; j < len(searchedString); j++ {
			if i+j >= len(aString) || aString[i+j] != searchedString[j] {
				found = false
				break
			}
		}
		if found {
			break
		}
	}
	if found {
		return i
	}
	return -1
}
