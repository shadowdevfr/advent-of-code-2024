package day02

import (
	"math"
	"strconv"
	"strings"
)

// Converts lines in format 7 6 4 2 1 to an array of int
func parseLine(line string) (result []int) {
	var unparsed []string = strings.Split(line, " ")
	for _, v := range unparsed {
		intval, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result = append(result, intval)
	}
	return result
}

// Checks if a given list is safe according to the puzzle
func isListSafe(tab []int) bool {
	var allIncreasing bool
	var init bool
	for i := 1; i < len(tab); i++ {
		var rate int = tab[i] - tab[i-1] // Decrease / Increase rate (< 0 decrease, >0 increase)
		if !init {
			allIncreasing = rate > 0
			init = true
		} else {
			if (rate > 0) != allIncreasing {
				// Not stable, unsafe
				return false
			}
		}

		if math.Abs(float64(rate)) > float64(3) || math.Abs(float64(rate)) < float64(1) {
			return false
		}
	}

	return true
}
