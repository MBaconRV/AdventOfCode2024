package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputAsRows := strings.Split(string(bytes), "\n")
	xmasCount := part1(inputAsRows)

	crossMasCount := part2(inputAsRows)

	fmt.Println("XMAS Count:", xmasCount)
	fmt.Println("X-MAS Count:", crossMasCount)
}

func part1(input []string) int {
	xmasCount := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'X' {
				if checkUp(input, i, j) {
					xmasCount++
				}
				if checkDown(input, i, j) {
					xmasCount++
				}
				if checkLeft(input, i, j) {
					xmasCount++
				}
				if checkRight(input, i, j) {
					xmasCount++
				}
				if checkUpRight(input, i, j) {
					xmasCount++
				}
				if checkUpLeft(input, i, j) {
					xmasCount++
				}
				if checkDownRight(input, i, j) {
					xmasCount++
				}
				if checkDownLeft(input, i, j) {
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}

func part2(input []string) int {
	crossMassCount := 0
	for i := range input {
		if i < 1 || i > len(input)-2 {
			continue
		}
		for j := range input[i] {
			if j < 1 || j > len(input[i])-2 {
				continue
			}
			if input[i][j] == 'A' {
				if (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M') {
					if (input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S') || (input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M') {
						crossMassCount++
					}
				}
			}
		}
	}
	return crossMassCount
}

func checkUp(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex < 3 {
		return false
	}
	return input[rowIndex-1][letterIndex] == 'M' && input[rowIndex-2][letterIndex] == 'A' && input[rowIndex-3][letterIndex] == 'S'
}

func checkDown(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex > len(input)-4 {
		return false
	}
	return input[rowIndex+1][letterIndex] == 'M' && input[rowIndex+2][letterIndex] == 'A' && input[rowIndex+3][letterIndex] == 'S'
}

func checkLeft(input []string, rowIndex int, letterIndex int) bool {
	if letterIndex < 3 {
		return false
	}
	return input[rowIndex][letterIndex-1] == 'M' && input[rowIndex][letterIndex-2] == 'A' && input[rowIndex][letterIndex-3] == 'S'
}

func checkRight(input []string, rowIndex int, letterIndex int) bool {
	if letterIndex > len(input[rowIndex])-4 {
		return false
	}
	return input[rowIndex][letterIndex+1] == 'M' && input[rowIndex][letterIndex+2] == 'A' && input[rowIndex][letterIndex+3] == 'S'
}

func checkUpRight(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex < 3 || letterIndex > len(input[rowIndex])-4 {
		return false
	}
	return input[rowIndex-1][letterIndex+1] == 'M' && input[rowIndex-2][letterIndex+2] == 'A' && input[rowIndex-3][letterIndex+3] == 'S'
}

func checkUpLeft(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex < 3 || letterIndex < 3 {
		return false
	}
	return input[rowIndex-1][letterIndex-1] == 'M' && input[rowIndex-2][letterIndex-2] == 'A' && input[rowIndex-3][letterIndex-3] == 'S'
}

func checkDownRight(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex > len(input)-4 || letterIndex > len(input[rowIndex])-4 {
		return false
	}
	return input[rowIndex+1][letterIndex+1] == 'M' && input[rowIndex+2][letterIndex+2] == 'A' && input[rowIndex+3][letterIndex+3] == 'S'
}

func checkDownLeft(input []string, rowIndex int, letterIndex int) bool {
	if rowIndex > len(input)-4 || letterIndex < 3 {
		return false
	}
	return input[rowIndex+1][letterIndex-1] == 'M' && input[rowIndex+2][letterIndex-2] == 'A' && input[rowIndex+3][letterIndex-3] == 'S'
}
