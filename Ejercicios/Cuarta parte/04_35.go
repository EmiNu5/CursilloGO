package main

import (
	"fmt"
	"math"
)

func main35() {
	var floats []float64 = []float64{}
	fmt.Println(len(floats))
	fmt.Println(minMaxFloat(floats))

}
func minMaxFloat(floats []float64) [2]float64 {
	var i int
	var max float64 = math.MinInt64
	var min float64 = math.MaxInt64
	var minMax [2]float64

	if len(floats) == 0 {
		max = 0
		min = 0
	} else {
		for i = 0; i < len(floats); i++ {
			if floats[i] > max {
				min = max
				max = floats[i]

			}
		}
	}
	minMax = [2]float64{max, min}
	return minMax
}
