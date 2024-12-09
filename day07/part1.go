package day07

import (
	"strconv"
	"strings"
)

type Operation struct {
	result   int
	operands []int
}

func parseLine(line string) Operation {
	var parts []string = strings.Split(line, ":")
	var operands []string = strings.Fields(parts[1])

	intres, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		panic(err)
	}

	var operandsfinal []int
	for _, op := range operands {
		opint, err := strconv.Atoi(op)
		if err != nil {
			panic(err)
		}
		operandsfinal = append(operandsfinal, opint)
	}

	return Operation{
		result:   intres,
		operands: operandsfinal,
	}
}

func generatePattern(n int) (patterns [][]string) {
	for i := 0; i < (1 << uint(n)); i++ {
		var pattern []string
		for j := 0; j < n; j++ {
			if (i & (1 << uint(j))) != 0 {
				pattern = append(pattern, "*")
			} else {
				pattern = append(pattern, "+")
			}
		}
		patterns = append(patterns, pattern)
	}
	return patterns
}

func solveEquation(op Operation) (solvingOperations []string) {
	var possibleOperations = generatePattern(len(op.operands) - 1)
	for _, operation := range possibleOperations {
		result := op.operands[0]
		for i := 0; i < len(operation); i++ {
			if operation[i] == "+" {
				result += op.operands[i+1]
			} else if operation[i] == "*" {
				result *= op.operands[i+1]
			}
		}
		if result == op.result {
			return operation
		}
	}
	return nil
}
