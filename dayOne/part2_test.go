package dayone

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPartTwoExample(t *testing.T) {
	if SolvePartTwo([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}) != 31 {
		t.Error()
	}
}

func TestPartTwoPuzzleInput(t *testing.T) {
	var left []int
	var right []int

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fail()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		leftn, rightn := lineToNumbers(scanner.Text())
		left = append(left, leftn)
		right = append(right, rightn)
	}

	var result int = SolvePartTwo(left, right)

	fmt.Printf("----- Solution is %d", result)

	if result != 23150395 {
		t.Fail()
	}
}
