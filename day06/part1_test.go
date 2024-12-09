package day06

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParseExample(t *testing.T) {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	var result [][]string = parseContent(string(content))
	var expected [][]string = [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error(result, expected)
	}
}

func TestGuardPos(t *testing.T) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	newPuzzleMap := DrawPath(parseContent(string(content)))
	var count int
	for _, l := range newPuzzleMap {
		for _, v := range l {
			if v == "X" {
				count++
			}
		}
	}

	fmt.Println("--------- Solution is", count)
	if count != 5030 {

	}
}
