package main

import (
	"fmt"
	"math"
)

func main() {
	sliceOfInt := []int{1, 1, 2, 3, 1, 5, 7, 3, 5, 4, 8, 90}
	sliceOfStr := []string{"apple", "banana", "apple", "orange", "banana", "banana"}

	mapA := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
		"peach":  4,
	}

	mapB := map[string]int{
		"apple":   3,
		"orange":  2,
		"tomatoe": 12,
		"coconut": 1,
	}

	rectangle := Rectangle{
		Width:  124,
		Height: 351,
	}

	circle := Cyrcle{
		Radius: 34,
	}

	fmt.Println(removeDuplicates(sliceOfInt))
	fmt.Println(countOccurrences(sliceOfStr))
	fmt.Println(mergeMaps(mapA, mapB))
	fmt.Println(RectArea(rectangle))
	fmt.Println(RectPerimeter(rectangle))
	fmt.Println(CircArea(circle))
	fmt.Println(Circumference(circle))
	fmt.Println("Yabadabadu")
	PrintShapeDetails(rectangle)
	PrintShapeDetails(circle)

}

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

func removeDuplicates(sliceOfInts []int) []int {

	newSliceOfInt := []int{}
	newSliceOfInt = append(newSliceOfInt, sliceOfInts[0])
	var isDuplicate bool

	for i := 1; i < len(sliceOfInts); i++ {
		isDuplicate = false
		for j := 0; j < len(newSliceOfInt); j++ {
			if sliceOfInts[i] == newSliceOfInt[j] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			newSliceOfInt = append(newSliceOfInt, sliceOfInts[i])
		}
	}
	return newSliceOfInt
}

func countOccurrences(sliceOfStrings []string) map[string]int {
	intOccMap := make(map[string]int)
	for i := 0; i < len(sliceOfStrings); i++ {
		_, isCont := intOccMap[sliceOfStrings[i]]
		if isCont {
			intOccMap[sliceOfStrings[i]] = intOccMap[sliceOfStrings[i]] + 1
		} else {
			intOccMap[sliceOfStrings[i]] = 1
		}
	}

	return intOccMap
}

func mergeMaps(mapA map[string]int, mapB map[string]int) map[string]int {
	for keyString := range mapB {
		_, isCont := mapA[keyString]
		if isCont {
			mapA[keyString] += mapB[keyString]
		} else {
			mapA[keyString] = mapB[keyString]
		}
	}
	return mapA
}
