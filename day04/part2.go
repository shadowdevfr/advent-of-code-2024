package day04

func SolvePartTwo(input [][]string) (amount int) {
	// navigate through the input with a margin of 1
	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if input[y][x] == "A" {
				if input[y-1][x-1] == "M" && input[y-1][x+1] == "S" && input[y+1][x+1] == "S" && input[y+1][x-1] == "M" {
					amount++
				} else if input[y-1][x-1] == "S" && input[y-1][x+1] == "M" && input[y+1][x+1] == "M" && input[y+1][x-1] == "S" {
					amount++
				} else if input[y-1][x-1] == "M" && input[y-1][x+1] == "M" && input[y+1][x+1] == "S" && input[y+1][x-1] == "S" {
					amount++
				} else if input[y-1][x-1] == "S" && input[y-1][x+1] == "S" && input[y+1][x+1] == "M" && input[y+1][x-1] == "M" {
					amount++
				}
			}
		}
	}

	return amount
}
