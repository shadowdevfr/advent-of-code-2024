package dayfive

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseSingleRule(t *testing.T) {
	var parsed Rule = parseRule("47|53")
	if parsed.first != 47 || parsed.second != 53 {
		t.Error(parsed)
	}
}

func TestParseSinglePageNumbersUpdate(t *testing.T) {
	var parsed []int = parsePageNumbersUpdates("75,47,61,53,29")
	var expected []int = []int{75, 47, 61, 53, 29}
	if !reflect.DeepEqual(parsed, expected) {
		t.Error(parsed)
	}
}

func TestRuleGood(t *testing.T) {
	var rule Rule = parseRule("47|53")
	var input []int = parsePageNumbersUpdates("75,47,61,53,29")

	if !isOrderGood([]Rule{rule}, input) {
		t.Fail()
	}
}

func TestRuleBad(t *testing.T) {
	var rule Rule = parseRule("53|47")
	var input []int = parsePageNumbersUpdates("75,47,61,53,29")

	if isOrderGood([]Rule{rule}, input) {
		t.Fail()
	}
}

func TestExample(t *testing.T) {

	file, err := os.Open("example.txt")
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
		if isOrderGood(rules, input) {
			var middle int = input[(len(input)-1)/2]
			sum += middle
			fmt.Println(input, middle)
		}
	}
	fmt.Printf("----- Solution is %d", sum)
	if sum != 143 {
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
		if isOrderGood(rules, input) {
			var middle int = input[(len(input)-1)/2]
			sum += middle
			fmt.Println(input, middle)
		}
	}
	fmt.Printf("----- Solution is %d", sum)
	if sum != 4185 {
		t.Fail()
	}
}
