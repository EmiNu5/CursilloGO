package main

import "fmt"

func main() {
	printAllCollors()
}

func printAllCollors() {
	const reset string = "\u001b[0m"
	var i int
	for i = 0; i < 256; i++ {
		if i%6 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s ", fmt.Sprintf("\u001b[38;5;%dm%3d", i, i))
	}
	fmt.Println(reset)
