package main

import (
	"fmt"
	"math"
	"strconv"
)

func GetTriangleInfo(sideA, sideB, sideC string) (string, []Coordinate) {
	a, err := strconv.ParseFloat(sideA, 64)
	if err != nil {
		return "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}
	}

	b, err := strconv.ParseFloat(sideB, 64)
	if err != nil {
		return "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}
	}

	c, err := strconv.ParseFloat(sideC, 64)
	if err != nil {
		return "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}
	}

	triangleType := GetTriangleType(a, b, c)
	triangleVertices := GetTriangleVertices(a, b, c)

	return triangleType, triangleVertices
}

func GetTriangleType(a, b, c float64) string {
	if a < 0 || b <= 0 || c < 0 {
		return ""
	} else if a+b <= c || a+c < b || b+c <= a {
		return "не треугольник"
	} else if a == b && b == c {
		return "равнобедренный"
	} else if a == b && b != c || a != c {
		return "равносторонний"
	} else {
		return "разносторонний"
	}
}

func GetTriangleVertices(a, b, c float64) []Coordinate {
	// A (угол между b и c)
	xA := 0
	yA := 0

	// B (угол между a и c)
	xB := int(c)
	yB := 0

	// C (угол между a и b)
	cosA := (b*b + c*c - a*a) / (2 * b * c)
	xC := int(b * cosA)
	yC := int(b * math.Sqrt(1-cosA*cosA))

	return []Coordinate{{xA, yA}, {xB, yB}, {xC, yC}}
}

func main() {
	triangleType, triangleVertices := GetTriangleInfo("3", "4", "5")
	fmt.Println("Тип треугольника:", triangleType)
	fmt.Println("Вершины треугольника:", triangleVertices)
}

type Coordinate struct {
	X, Y int
}
