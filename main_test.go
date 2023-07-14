package main

import (
	"reflect"
	"testing"
)

// Test RemoveDuplicates

func TestRemoveDuplicates(t *testing.T) {
	sliceOfInts := []int{1, 2, 2, 3, 4, 4, 5, 5, 5}

	expected := []int{1, 2, 3, 4, 5}
	actual := removeDuplicates(sliceOfInts)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestRemoveDuplicates_EmptySlice(t *testing.T) {
	sliceOfInts := []int{}

	expected := []int{}
	actual := removeDuplicates(sliceOfInts)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestRemoveDuplicates_NoDuplicates(t *testing.T) {
	sliceOfInts := []int{1, 2, 3, 4, 5}

	expected := []int{1, 2, 3, 4, 5}
	actual := removeDuplicates(sliceOfInts)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

// Test COuntOccurrences

func TestCountOccurrences(t *testing.T) {
	sliceOfStrings := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

	expected := map[string]int{
		"apple":  3,
		"banana": 2,
		"orange": 1,
	}
	actual := countOccurrences(sliceOfStrings)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestCountOccurrences_EmptySlice(t *testing.T) {
	sliceOfStrings := []string{}

	expected := map[string]int{}
	actual := countOccurrences(sliceOfStrings)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestCountOccurrences_OneElement(t *testing.T) {
	sliceOfStrings := []string{"apple"}

	expected := map[string]int{
		"apple": 1,
	}
	actual := countOccurrences(sliceOfStrings)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

//Test MergeMaps

func TestMergeMaps(t *testing.T) {
	mapA := map[string]int{"a": 1, "b": 2}
	mapB := map[string]int{"b": 3, "c": 4}

	expected := map[string]int{"a": 1, "b": 5, "c": 4}
	actual := mergeMaps(mapA, mapB)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestMergeMaps_EmptyMaps(t *testing.T) {
	mapA := map[string]int{}
	mapB := map[string]int{}

	expected := map[string]int{}
	actual := mergeMaps(mapA, mapB)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestMergeMaps_MapBEmpty(t *testing.T) {
	mapA := map[string]int{"a": 1, "b": 2}
	mapB := map[string]int{}

	expected := map[string]int{"a": 1, "b": 2}
	actual := mergeMaps(mapA, mapB)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestMergeMaps_MapAEmpty(t *testing.T) {
	mapA := map[string]int{}
	mapB := map[string]int{"a": 1, "b": 2}

	expected := map[string]int{"a": 1, "b": 2}
	actual := mergeMaps(mapA, mapB)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}
}

/// Test ReverseString

func TestReverseString(t *testing.T) {
	str := "Hello, World!"

	expected := "!dlroW ,olleH"
	actual := ReverseString(str)

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %s, Actual: %s", expected, actual)
	}
}

func TestReverseString_EmptyString(t *testing.T) {
	str := ""

	expected := ""
	actual := ReverseString(str)

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %s, Actual: %s", expected, actual)
	}
}

func TestReverseString_SingleCharacter(t *testing.T) {
	str := "A"

	expected := "A"
	actual := ReverseString(str)

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %s, Actual: %s", expected, actual)
	}
}
