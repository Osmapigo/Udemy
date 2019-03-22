package main

import "testing"

func TestAverage(t *testing.T) {
	dataPoints := []int{1, 2, 3, 4, 4}
	result := average(dataPoints)
	if result != 2.8 {
		t.Error("Expected value 2.8, obtained: ", result)
	}
}
