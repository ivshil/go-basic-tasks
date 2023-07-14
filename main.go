package main

import (
	"fmt"
	"go-basic-tasks/src/models"
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

	rectangle := models.Rectangle{
		Width:  124,
		Height: 351,
	}

	circle := models.Cyrcle{
		Radius: 34,
	}

	fmt.Println(removeDuplicates(sliceOfInt))
	fmt.Println(countOccurrences(sliceOfStr))
	fmt.Println(mergeMaps(mapA, mapB))
	fmt.Println(models.RectArea(rectangle))
	fmt.Println(models.RectPerimeter(rectangle))
	fmt.Println(models.CircArea(circle))
	fmt.Println(models.Circumference(circle))
	fmt.Println("Yabadabadu")
	models.PrintShapeDetails(rectangle)
	models.PrintShapeDetails(circle)

	name := "Gencho"
	fmt.Println(ReverseString(name))

	person := models.Person{
		Name:  "John Doe",
		Age:   25,
		Email: "johndoe@example.org",
	}

	jsonPerson := models.EncodeJSON(person)
	fmt.Println(jsonPerson)
	fmt.Println(models.DecodeJSON(jsonPerson))

}

func ReverseString(str string) string {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string([]rune(str)[i])
	}
	return newStr
}

func removeDuplicates(sliceOfInts []int) []int {

	if len(sliceOfInts) == 0 {
		return []int{}
	}

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
