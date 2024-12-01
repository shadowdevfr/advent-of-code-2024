package dayone

import (
	"math"
	"strconv"
	"strings"
)

const separator = "   " // This separates the two lists of the puzzle input

// This function takes a single line from the puzzle input and transforms it into two numbers left and right.
func lineToNumbers(line string) (left int, right int) {
	var parts []string = strings.Split(line, separator)

	// Parse the ints
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	right, err = strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return left, right
}

// Implements the merge sort algorithm for efficient sorting that will be useful later on
func mergeSort(tab []int) []int {
	if len(tab) == 1 {
		return tab
	}

	var millieu int = len(tab) / 2
	var left []int = tab[:millieu]
	var right []int = tab[millieu:]

	left = mergeSort(left)
	right = mergeSort(right)

	tab = make([]int, len(tab))

	var gi int
	var di int
	var i int
	for gi < len(left) && di < len(right) {
		if left[gi] < right[di] {
			tab[i] = left[gi]
			gi++
		} else {
			tab[i] = right[di]
			di++
		}
		i++
	}

	// On ajoute les éléments restants
	for gi < len(left) {
		tab[i] = left[gi]
		gi++
		i++
	}
	for di < len(right) {
		tab[i] = right[di]
		di++
		i++
	}

	return tab
}

// Solves the puzzle with two given lists left and right
func Solve(left []int, right []int) (sum int) {
	var sortedLeft []int = mergeSort(left)
	var sortedRight []int = mergeSort(right)

	// the two lists have the same length as provided in the puzzle input
	for i := 0; i < len(sortedLeft); i++ {
		sum += int(math.Abs(float64(sortedRight[i] - sortedLeft[i])))
	}

	return sum
}
