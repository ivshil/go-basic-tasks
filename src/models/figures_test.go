package models

import (
	"math"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 5}
	expected := 20.0
	actual := rect.Area()

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %f, Actual: %f", expected, actual)
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 5}
	expected := 18.0
	actual := rect.Perimeter()

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %f, Actual: %f", expected, actual)
	}
}

func TestCyrcle_Area(t *testing.T) {
	circle := Cyrcle{Radius: 3}
	expected := math.Pi * 9.0
	actual := circle.Area()

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %f, Actual: %f", expected, actual)
	}
}

func TestCyrcle_Perimeter(t *testing.T) {
	circle := Cyrcle{Radius: 3}
	expected := math.Pi * 6.0
	actual := circle.Perimeter()

	if actual != expected {
		t.Errorf("Unexpected result. Expected: %f, Actual: %f", expected, actual)
	}
}
