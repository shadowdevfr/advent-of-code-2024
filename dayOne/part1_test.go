package dayone

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPossibleGame(t *testing.T) {
	var game Game = Game{sets: []Set{
		{blue: 3, red: 4},
		{blue: 6, red: 1, green: 2},
		{green: 2},
	}}

	if !game.isPossible(Set{red: 12, green: 13, blue: 14}) {
		t.Fail()
	}
}

func TestImpossibleGame(t *testing.T) {
	var game Game = Game{sets: []Set{
		{blue: 3, red: 4},
		{blue: 6, red: 1, green: 2},
		{green: 2},
	}}

	if game.isPossible(Set{red: 1, green: 2, blue: 3}) {
		t.Fail()
	}
}

func TestConvertLineToGame(t *testing.T) {
	var result Game = convertGamestringLineIntoGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	var valeursouhaitee []Set = []Set{{red: 4, blue: 3}, {red: 1, green: 2, blue: 6}, {green: 2}}
	if !reflect.DeepEqual(result.sets, valeursouhaitee) {
		t.Fail()
	}
}

func TestSolveExample(t *testing.T) {
	var sum int = Solve("example.txt")
	if sum != 8 {
		t.Error("Puzzle returned", sum, "instead of 8")
	}
}

func TestSolvePuzzle(t *testing.T) {
	var sum int = Solve("input.txt")
	fmt.Println("Solution to the puzzle is", sum)
	if sum != 2439 {
		t.Error("Puzzle returned", sum, "instead of 2439")
	}
}
