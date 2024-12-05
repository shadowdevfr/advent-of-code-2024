package day02

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPuzzleInputStep2(t *testing.T) {

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fail()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		if isListSafeWithDampener(parseLine(scanner.Text())) {
			result++
		}
	}

	fmt.Printf("----- Solution is %d", result)

	if result != 589 {
		t.Error()
	}
}
