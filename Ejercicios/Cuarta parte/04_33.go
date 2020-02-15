package main

import "fmt"

func main33() {
	var floats []float64 = []float64{1.5, 2.3, 5.2}

	fmt.Println(sumFloat(floats))
}
func sumFloat(floats []float64) float64 {
	var i int
	var sum float64

	for i = 0; i < len(floats); i++ {
		sum = sum + floats[i]
	}
	return sum
}
