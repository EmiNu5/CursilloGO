package main

import "fmt"

func main() {
	var unString string = "HELLO PEOPLE"
	inRed(unString)
}
func inRed(unString string) {
	const inRed string = "\033[1;31m%s\033[0m"
	fmt.Printf(inRed, unString)
}
