package day03

import (
	"fmt"
	"os"
	"testing"
)

func TestPuzzleInputPartTwo(t *testing.T) {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		t.Error(err)
	}

	var result int = SolvePartTwo(string(file))

	fmt.Printf("----- Solution is %d", result)

	if result != 89823704 {
		t.Fail()
	}
}
