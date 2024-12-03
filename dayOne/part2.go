package dayone

func SolvePartTwo(left []int, right []int) (similarityScore int) {
	for _, x := range left {
		var amount int
		for _, y := range right {
			if y == x {
				amount++
			}
		}
		similarityScore += amount * x
	}
	return similarityScore
}
