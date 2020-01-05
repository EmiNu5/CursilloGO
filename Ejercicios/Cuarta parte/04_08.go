package main

import "fmt"

func main() {
	var unString string = "QUE ONDA LA PEOPLE"
	inYelow(unString)
}
func inYelow(unString string) {
	const yellow string = " \u001b[33m"
	fmt.Print(yellow, unString)
}
