package day09

import (
	"strconv"
)

func parseDiskMap(input string) ([]int, []int) {
	files := []int{}
	freeSpaces := []int{}
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(string(input[i]))
		if i%2 == 0 {
			files = append(files, n)
		} else {
			freeSpaces = append(freeSpaces, n)
		}
	}
	return files, freeSpaces
}

func compactDisk(files []int, freeSpaces []int) []int {
	disk := []int{}
	for i, file := range files {
		for j := 0; j < file; j++ {
			disk = append(disk, i)
		}
		if i < len(freeSpaces) {
			for j := 0; j < freeSpaces[i]; j++ {
				disk = append(disk, -1) // -1 represents free space
			}
		}
	}

	for i, e := range disk {
		if e == -1 {
			for j := len(disk) - 1; j > i; j-- {
				if disk[j] != -1 {
					disk[j], disk[i] = disk[i], disk[j]
					break
				}
			}
		}
	}

	return disk
}

func calculateChecksum(disk []int) int {
	checksum := 0
	for i, value := range disk {
		if value != -1 {
			checksum += i * value
		}
	}
	return checksum
}
