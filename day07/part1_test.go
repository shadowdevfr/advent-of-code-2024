package day07

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var expected Operation = Operation{
		result:   190,
		operands: []int{10, 19},
	}

	var result Operation = parseLine("190: 10 19")

	if expected.result != result.result || !reflect.DeepEqual(expected.operands, result.operands) {
		t.Fail()
	}
}

func TestSolveEquation(t *testing.T) {
	var result []string = solveEquation(Operation{result: 190, operands: []int{10, 19}})

	if result[0] != "*" {
		t.Fail()
	}
}

func TestSolveEquationUnsolvable(t *testing.T) {
	var result []string = solveEquation(Operation{result: 21037, operands: []int{9, 7, 18, 13}})

	if len(result) != 0 {
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
	var result int
	for scanner.Scan() {
		var op Operation = parseLine(scanner.Text())
		var res []string = solveEquation(op)
		if len(res) != 0 {
			result += op.result
		}
	}

	fmt.Printf("----- Solution is %d", result)

	if result != 12940396350192 {
		t.Error()
	}
}
