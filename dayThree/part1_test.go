package dayThree

import (
	"fmt"
	"os"
	"testing"
)

func TestValidMul(t *testing.T) {
	if !isValidMul("mul(1,2)") {
		t.Error()
	}
}

func TestInvalidMul(t *testing.T) {
	if isValidMul("mul[1,2!") {
		t.Error()
	}
}

func TestParseMul(t *testing.T) {
	left, right := parseMul("mul(1,4)")
	if left != 1 && right != 4 {
		t.Error()
	}
}

func TestExampleSolve(t *testing.T) {
	if Solve("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))") != 161 {
		t.Error()
	}
}

func TestPuzzleInput(t *testing.T) {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		t.Error(err)
	}

	var result int = Solve(string(file))

	fmt.Printf("----- Solution is %d", result)

	if result != 167090022 {
		t.Fail()
	}
}
