package dayThree

// Takes a mul (AND A MUL ONLY) as the input and returns true if it's valid.
const do string = "do()"
const dont string = "don't()"

func SolvePartTwo(puzzle string) (result int) {
	var activated bool = true
	for i := 0; i < len(puzzle)-2; i++ {
		// do() / don't() detector
		if string(puzzle[i]) == "d" && string(puzzle[i+1]) == "o" {
			// detect if it's a do or don't
			if i+len(dont) < len(puzzle) {
				if puzzle[i:i+len(dont)] == dont {
					activated = false
					continue
				}
			}
			if i+len(do) < len(puzzle) {
				if puzzle[i:i+len(do)] == do {
					activated = true
				}
			}
		}
		// mul detector
		if activated && string(puzzle[i]) == "m" && string(puzzle[i+1]) == "u" && string(puzzle[i+2]) == "l" {
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
