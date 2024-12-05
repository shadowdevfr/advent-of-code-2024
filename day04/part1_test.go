package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	var result int = Solve([][]string{{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	})

	if result != 18 {
		t.Fail()
	}
}

func TestPuzzleInput(t *testing.T) {

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fail()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]string
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	var result int = Solve(input)

	fmt.Printf("----- Solution is %d", result)

	if result != 2468 {
		t.Error()
	}
}
