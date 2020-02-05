package main

import "fmt"

func main8() {
	var unString string = "QUE ONDA LA PEOPLE"
	fmt.Print(inYelow(unString))
}
func inYelow(unString string) string {
	const yellow string = " \u001b[33m"
	const reinicio string = "\u001b[0m"
	return yellow + unString + reinicio
}
