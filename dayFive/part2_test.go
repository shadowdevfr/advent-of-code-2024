package dayfive

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPuzzleInputStepTwo(t *testing.T) {

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fail()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules []Rule
	var inputs [][]int
	for scanner.Scan() {
		var txt string = scanner.Text()
		if strings.Contains(txt, "|") {
			rules = append(rules, parseRule(txt))
		} else if strings.Contains(txt, ",") {
			inputs = append(inputs, parsePageNumbersUpdates(txt))
		}
	}
	var sum int
	for _, input := range inputs {
		if !isOrderGood(rules, input) {
			orderedUpdate := fixOrder(rules, input)
			middleIndex := len(orderedUpdate) / 2
			sum += orderedUpdate[middleIndex]
		}
	}
	fmt.Printf("----- Solution is %d", sum)
	if sum != 4480 {
		t.Fail()
	}
}
