package main

import "fmt"

func main() {
	dataPoints := []int{1, 2, 3, 4, 4}
	fmt.Println(average(dataPoints))
}

func average(points []int) float64 {
	val := 0
	len := len(points)
	for _, number := range points {
		val = val + number
	}
	return float64(val) / float64(len)
}
