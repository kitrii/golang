package main

import (
	"fmt"
	"math"
)

func main() {
	object := Circle{10, 20, 30}
	fmt.Println(object)
	fmt.Println(CircleArea(object))
	fmt.Println(object)
	fmt.Println(object.area())
	fmt.Println(object)

}

type Circle struct {
	x, y, r float64
}

func CircleArea(c Circle) float64 {
	result := c.r * c.r * math.Pi
	c.x += 100
	return result

}

func (c *Circle) area() float64 {
	result := c.r * c.r * math.Pi
	c.y += 100
	return result

}
