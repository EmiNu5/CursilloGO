package main

import "fmt"

func main6() {
	var unString string = "HELLO PEOPLE"
	fmt.Println(inGreen(unString))
}
func inGreen(unString string) string {
	const green string = "\u001b[32m"
	const reinicio string = "\u001b[0m"

	return green + unString + reinicio
}
