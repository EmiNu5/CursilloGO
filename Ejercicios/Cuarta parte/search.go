package main

import "fmt"

func main() {
	var str1 = "papa"
	var str2 = "apa"

	var i int
	var j int

	var count int = 1
	var ocurrence []int
	var aux []int

	for i = 0; i < len(str1); i++ {
		for j = 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				count++
				aux = append(ocurrence, i)
				if count == len(str2) {
					fmt.Println(count, "==", len(str2))
					break
				}
			} else {
				break
			}

		}
	}
	fmt.Println(aux[0])
}
