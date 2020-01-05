package main

import "fmt"

func main() {
	var unString string = "HELLO PEOPLE"
	InGreen(unString)
}
func InGreen(unString string) {
	const InGreen string = "\033[1;32m%s\033[0m"
	fmt.Printf(InGreen, unString)
}
