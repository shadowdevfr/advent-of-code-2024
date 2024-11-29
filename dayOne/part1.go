package dayone

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int // number of red cubes
	green int // number of green cubes
	blue  int // number of blue cubes
}

// Determines if the game is possible given a number of r g and b cubes
func (g *Game) isPossible(set Set) bool {
	for _, s := range g.sets {
		if s.blue > set.blue || s.green > set.green || s.red > set.red {
			return false
		}
	}
	return true
}

func convertGamestringLineIntoGame(str string) (g Game) {
	var splitted []string = strings.Split(str, ":")
	var sets []string = strings.Split(splitted[1], ";")
	gameId, err := strconv.Atoi(strings.Split(splitted[0], " ")[1])
	if err != nil {
		panic(err)
	}
	g.id = gameId
	for _, s := range sets {
		var cubes []string = strings.Split(s, ",")
		var set Set = Set{}
		for _, c := range cubes {
			var cubeValues []string = strings.Split(c, " ")
			nbCubes, err := strconv.Atoi(cubeValues[1])
			if err != nil {
				continue
			}
			var cubeColor string = cubeValues[2]

			if cubeColor == "red" {
				set.red += nbCubes
			}
			if cubeColor == "green" {
				set.green += nbCubes
			}
			if cubeColor == "blue" {
				set.blue += nbCubes
			}
		}
		g.sets = append(g.sets, set)
	}
	return g
}

func Solve(filename string) (sum int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var set Set = Set{red: 12, green: 13, blue: 14}
	for scanner.Scan() {
		var game Game = convertGamestringLineIntoGame(scanner.Text())
		if game.isPossible(set) {
			sum += game.id
		}
	}
	return sum
}
