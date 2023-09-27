package main

import (
	"fmt"
	"strconv"
)

func CalculateKindTriangle(lengthA string, lengthB string, lengthC string) (kindTriangle string, coordinate [3][2]int) {
	floatLengthA, err := strconv.ParseFloat(lengthA, 64)
	floatLengthB, err := strconv.ParseFloat(lengthB, 64)
	floatLengthC, err := strconv.ParseFloat(lengthC, 64)

	if floatLengthA == floatLengthB && floatLengthB == floatLengthC {
		return "равнобедренный", coordinate
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
