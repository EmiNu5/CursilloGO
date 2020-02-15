package main

import "fmt"

func main34() {
	var floats []float64 = []float64{1.5, 2.3, 5.2}

	fmt.Println(averageFloat(floats))
}
func averageFloat(floats []float64) float64 {
	var prom float64
	prom = sumFloat(floats) / float64(len(floats))

	return prom
}
