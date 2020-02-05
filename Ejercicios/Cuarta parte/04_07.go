package main

import "fmt"

func main7() {
	var unString string = "THE MANDALORIAN"
	fmt.Print(inRed(unString))
}
func inRed(unString string) string {
	const inRed string = "\u001b[31m"
	const reinicio string = "\u001b[0m"
	return inRed + unString + reinicio
}
