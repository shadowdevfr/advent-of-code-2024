package dayThree

import (
	"strconv"
	"strings"
)

// Takes a mul (AND A MUL ONLY) as the input and returns true if it's valid.
const allowedChars = "mul()1234567890,"

func isValidMul(str string) bool {
	var criterias []int = make([]int, len(allowedChars)) // connected to allowedChars, a mul must have {1,1,1,1,1,x,x,x,x,x,x,x,x,x,x,1}
	for _, c := range str {
		for aci, ac := range allowedChars {
			if ac == c {
				criterias[aci] += 1
			}
		}
	}

	return criterias[0] == 1 && criterias[1] == 1 && criterias[2] == 1 && criterias[3] == 1 && criterias[4] == 1 && criterias[len(criterias)-1] == 1
}

func parseMul(mul string) (left int, right int) {
	var first string = strings.Replace(mul, "mul(", "", 1)
	var second string = strings.Replace(first, ")", "", 1)
	var split []string = strings.Split(second, ",")
	left, err := strconv.Atoi(split[0])
	if err != nil {
		return
	}
	right, err = strconv.Atoi(split[1])
	if err != nil {
		return
	}
	return left, right
}

func Solve(puzzle string) (result int) {
	for i := 0; i < len(puzzle)-2; i++ {
		if string(puzzle[i]) == "m" && string(puzzle[i+1]) == "u" && string(puzzle[i+2]) == "l" {
			var mul string
			for j := i; j < len(puzzle); j++ {
				var found bool
				for _, ac := range allowedChars {
					if byte(ac) == puzzle[j] {
						found = true
						mul += string(ac)
					}
					if string(puzzle[j]) == ")" { // doing this so connected muls can be delimited
						mul += ")"
						break
					}
				}
				if !found {
					break
				}

			}
			if isValidMul(mul) {
				left, right := parseMul(mul)
				result += left * right
			}
		}
	}
	return result
}
