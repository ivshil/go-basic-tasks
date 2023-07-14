package models

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeDetails(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func RectArea(r Rectangle) float64 {
	return r.Width * r.Height
}

func RectPerimeter(r Rectangle) float64 {
	return r.Width*2 + r.Height*2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Height*2
}

type Cyrcle struct {
	Radius float64
}

func CircArea(c Cyrcle) float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func Circumference(c Cyrcle) float64 {
	return (c.Radius * 2) * math.Pi
}

func (c Cyrcle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Cyrcle) Perimeter() float64 {
	return (c.Radius * 2) * math.Pi
}
