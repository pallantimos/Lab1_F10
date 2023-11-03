package main

import (
	"reflect"
	"testing"
)

func TestGetTriangleInfo(t *testing.T) {
	tests := []struct {
		sideA      string
		sideB      string
		sideC      string
		wantType   string
		wantCoords []Coordinate
	}{
		// Тесты для правильных значений сторон треугольника
		{"3", "4", "5", "разносторонний", []Coordinate{{0, 0}, {5, 0}, {4, 3}}},
		{"4", "4", "4", "равнобедренный", []Coordinate{{0, 0}, {4, 0}, {2, 3}}},
		{"5", "5", "5", "равносторонний", []Coordinate{{0, 0}, {5, 0}, {2, 4}}},
		{"5", "12", "13", "разносторонний", []Coordinate{{0, 0}, {13, 0}, {5, 12}}},
		{"8", "15", "17", "разносторонний", []Coordinate{{0, 0}, {17, 0}, {8, 15}}},

		// Тесты для неправильных значений сторон треугольника
		{"0", "0", "0", "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}},
		{"-2", "3", "4", "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}},
		{"5", "10", "25", "не треугольник", []Coordinate{{-4, -1}, {-1, -1}, {-1, -1}}},

		// Тесты для некорректных входных данных
		{"abc", "def", "ghi", "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}},
		{"1", "", "3", "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}},
		{"1", "2", "3", "", []Coordinate{{-2, -2}, {-2, -2}, {-4, -2}}},
	}

	for _, test := range tests {
		gotType, gotCoords := GetTriangleInfo(test.sideA, test.sideB, test.sideC)
		if gotType != test.wantType {
			t.Errorf("GetTriangleInfo(%s, %s, %s) = %s; want %s", test.sideA, test.sideB, test.sideC, gotType, test.wantType)
		}
		if !reflect.DeepEqual(gotCoords, test.wantCoords) {
			t.Errorf("GetTriangleInfo(%s, %s, %s) coordinates = %v; want %v", test.sideA, test.sideB, test.sideC, gotCoords, test.wantCoords)
		}
	}
}
