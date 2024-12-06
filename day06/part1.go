package day06

import (
	"bufio"
	"fmt"
	"strings"
)

// Parses input to an array of string arrays
func parseContent(content string) (result [][]string) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}
	return result
}

func DrawPath(puzzleMap [][]string) (newPuzzleMap [][]string) {
	var guardX int
	var guardY int
	var guardDirection string

	// Find current guard x and y
	for y := 0; y < len(puzzleMap); y++ {
		for x := 0; x < len(puzzleMap[y]); x++ {
			if puzzleMap[y][x] == "^" || puzzleMap[y][x] == ">" || puzzleMap[y][x] == "v" || puzzleMap[y][x] == "<" {
				guardX = x
				guardY = y
				guardDirection = puzzleMap[y][x]
			}
		}
	}

	for {
		if guardDirection == "^" {
			for guardY >= 0 && puzzleMap[guardY][guardX] != "#" {
				fmt.Println(guardDirection, guardX, guardY)
				puzzleMap[guardY][guardX] = "X"
				guardY--
			}
			if guardY < 0 {
				return puzzleMap
			}
			guardDirection = ">"
			guardY++
			fmt.Println("switch to", guardDirection, guardX, guardY)
		}
		if guardDirection == "v" {
			for guardY < len(puzzleMap) && puzzleMap[guardY][guardX] != "#" {
				fmt.Println(guardDirection, guardX, guardY)
				puzzleMap[guardY][guardX] = "X"
				guardY++
			}
			if guardY >= len(puzzleMap) {
				return puzzleMap
			}
			guardDirection = "<"
			guardY--
			fmt.Println("switch to", guardDirection, guardX, guardY)
		}
		if guardDirection == ">" {
			for guardX < len(puzzleMap[guardY]) && puzzleMap[guardY][guardX] != "#" {
				fmt.Println(guardDirection, guardX, guardY)
				puzzleMap[guardY][guardX] = "X"
				guardX++
			}
			if guardX >= len(puzzleMap[guardY]) {
				return puzzleMap
			}
			guardDirection = "v"
			guardX--
			fmt.Println("switch to", guardDirection, guardX, guardY)
		}
		if guardDirection == "<" {
			for guardX >= 0 && puzzleMap[guardY][guardX] != "#" {
				fmt.Println(guardDirection, guardX, guardY)
				puzzleMap[guardY][guardX] = "X"
				guardX--
			}
			if guardX < 0 {
				return puzzleMap
			}
			guardDirection = "^"
			guardX++
			fmt.Println("switch to", guardDirection, guardX, guardY)
		}
	}
}
