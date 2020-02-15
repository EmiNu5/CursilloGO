package main

import "fmt"

func main30() {

	var strings []string = []string{"hola", "mundo", "!"}
	var lengths []int = mapToLength(strings)
	fmt.Println(lengths)

}
func mapToLength(strings []string) []int {
	var i int
	var lenth []int

	for i = 0; i < len(strings); i++ {
		lenth = append(lenth, len(strings[i]))
	}
	return lenth
}
