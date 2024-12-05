package day01

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLineToNumbers(t *testing.T) {
	left, right := lineToNumbers("3   4")

	if left != 3 || right != 4 {
		t.Errorf("lineToNumbers returned %d, %d instead of 3, 4", left, right)
	}
}

func TestMergeSort(t *testing.T) {
	var input []int = []int{19, 4, 8, 6, 9, 8, 12, 43, 30}
	var result []int = mergeSort(input)
	var expected []int = []int{4, 6, 8, 8, 9, 12, 19, 30, 43}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestExample(t *testing.T) {
	var left []int = []int{3, 4, 2, 1, 3, 3}
	var right []int = []int{4, 3, 5, 3, 9, 3}
	if Solve(left, right) != 11 {
		t.Fail()
	}
}

func TestPuzzleInput(t *testing.T) {
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

	var result int = Solve(left, right)

	fmt.Printf("----- Solution is %d", result)

	if result != 936063 {
		t.Fail()
	}
}
