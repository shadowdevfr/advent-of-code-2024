package day09

import (
	"fmt"
	"os"
	"testing"
)

func TestSolvePuzzleInput(t *testing.T) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var checksum int = calculateChecksum(compactDisk(parseDiskMap(string(content))))

	fmt.Println("--------- Solution is", checksum)
	if checksum != 5030 {

	}
}
