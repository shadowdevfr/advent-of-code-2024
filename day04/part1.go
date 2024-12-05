package day04

import (
	"reflect"
)

var xmas []string = []string{"X", "M", "A", "S"}
var invxmas []string = []string{"S", "A", "M", "X"}

func Solve(input [][]string) (amount int) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			// Check horizontally order and inverted
			if x+len(xmas) <= len(input[y]) {
				if reflect.DeepEqual(input[y][x:x+len(xmas)], xmas) {
					amount++
				} else if reflect.DeepEqual(input[y][x:x+len(xmas)], invxmas) {
					amount++
				}
			}
		}
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			// Check vertically order and inverted
			if y+len(xmas) <= len(input) {
				verticalSlice := make([]string, len(xmas))
				for i := 0; i < len(xmas); i++ {
					verticalSlice[i] = input[y+i][x]
				}
				if reflect.DeepEqual(verticalSlice, xmas) {
					amount++
				} else if reflect.DeepEqual(verticalSlice, invxmas) {
					amount++
				}
			}
		}
	}

	// Check diagonally (top-left to bottom-right) order and inverted
	for y := 0; y < len(input)-len(xmas)+1; y++ {
		for x := 0; x < len(input[y])-len(xmas)+1; x++ {
			diagonalSlice := make([]string, len(xmas))
			for i := 0; i < len(xmas); i++ {
				diagonalSlice[i] = input[y+i][x+i]
			}
			if reflect.DeepEqual(diagonalSlice, xmas) {
				amount++
			} else if reflect.DeepEqual(diagonalSlice, invxmas) {
				amount++
			}
		}
	}

	// Check diagonally (top-right to bottom-left) order and inverted
	for y := 0; y < len(input)-len(xmas)+1; y++ {
		for x := len(xmas) - 1; x < len(input[y]); x++ {
			diagonalSlice := make([]string, len(xmas))
			for i := 0; i < len(xmas); i++ {
				diagonalSlice[i] = input[y+i][x-i]
			}
			if reflect.DeepEqual(diagonalSlice, xmas) {
				amount++
			} else if reflect.DeepEqual(diagonalSlice, invxmas) {
				amount++
			}
		}
	}

	return amount
}
