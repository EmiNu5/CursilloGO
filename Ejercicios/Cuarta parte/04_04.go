package main

import "fmt"

func main4() {
	var unString string = "hola ffffff"
	printTitle(unString)
}
func printTitle(unString string) {

	var i int

	fmt.Printf("%s\n", unString)
	for i = 0; i < len(unString); i++ {
		if unString[i] != ' ' {
			fmt.Print("-")
		} else {
			fmt.Print(" ")
		}
	}

}
