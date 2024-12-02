package dayTwo

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var expected []int = []int{7, 6, 4, 2, 1}
	var result []int = parseLine("7 6 4 2 1")

	if !reflect.DeepEqual(result, expected) {
		t.Error()
	}
}

func TestIsListSafeSafe(t *testing.T) {
	if !isListSafe([]int{7, 6, 4, 2, 1}) {
		t.Error()
	}
}

func TestIsListSafeNotSafe1(t *testing.T) {
	if isListSafe([]int{1, 2, 7, 8, 9}) { // Rate over 3 or under 1
		t.Error()
	}
	if isListSafe([]int{1, 3, 2, 4, 5}) { // Not all increasing / decreasing
		t.Error()
	}
}

func TestPuzzleInput(t *testing.T) {

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fail()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		if isListSafe(parseLine(scanner.Text())) {
			result++
		}
	}

	fmt.Printf("----- Solution is %d", result)

	if result != 549 {
		t.Error()
	}
}
