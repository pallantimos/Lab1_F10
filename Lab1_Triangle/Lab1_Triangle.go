package main

import (
	"fmt"
	"strconv"
)

func CalculateKindTriangle(lengthA string, lengthB string, lengthC string) (kindTriangle string, coordinate [3][2]int) {
	floatLengthA, err := strconv.ParseFloat(lengthA, 64)
	floatLengthB, err := strconv.ParseFloat(lengthB, 64)
	floatLengthC, err := strconv.ParseFloat(lengthC, 64)

	if floatLengthA+floatLengthB > floatLengthC || floatLengthA+floatLengthC > floatLengthB || floatLengthB+floatLengthC > floatLengthA {
		return "не треугольник", coordinate
	}
	if floatLengthA == floatLengthB && floatLengthB == floatLengthC {
		return "равносторонний", coordinate
	} else if floatLengthA == floatLengthB || floatLengthB == floatLengthC || floatLengthC == floatLengthA {
		coordinate := [3][2]int{
			{0, 0},
			{50, 100},
			{100, 0},
		}
		return "равнобедренный", coordinate
	} else {
		return "разносторонний", coordinate
	}

	if err != nil {
		fmt.Println("Ошибка при конвертации", err)
		coordinate := [3][2]int{
			{-2, -2},
			{-2, -2},
			{-2, -2},
		}
		return "", coordinate
	}
	return kindTriangle, coordinate
}
