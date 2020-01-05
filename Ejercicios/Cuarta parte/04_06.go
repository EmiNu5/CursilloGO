package main

import "fmt"

func main() {
	var unString string = "HELLO PEOPLE"
	inGreen(unString)
}
func inGreen(unString string) {
	const InGreen string = "\033[1;32m%s\033[0m"
	fmt.Printf(InGreen, unString)
}
