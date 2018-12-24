package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return 2 * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		radius: 12.1,
	}
	sq1 := square{
		length: 22,
		width:  2.2,
	}
	info(c1)
	info(sq1)
}
