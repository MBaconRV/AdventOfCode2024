package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)

	formattedInput := formatInput(input)

	safetyCount := 0
	safetyCountWithDeletion := 0

	for _, row := range formattedInput {
		if checkAscending(row) {
			safetyCount++
			safetyCountWithDeletion++
		} else if checkDescending(row) {
			safetyCount++
			safetyCountWithDeletion++
		} else {
			for i := 0; i < len(row); i++ {
				if checkAscending(remove(row, i)) {
					safetyCountWithDeletion++
					break
				} else if checkDescending(remove(row, i)) {
					safetyCountWithDeletion++
					break
				}
			}
		}
	}

	fmt.Println("safety count:", safetyCount)
	fmt.Println("safety count with deletions:", safetyCountWithDeletion)
}

func formatInput(input string) (formattedInput [][]int) {
	inputAsRows := strings.Split(input, "\n")

	for i := 0; i < len(inputAsRows); i++ {
		stringNumbers := strings.Split(inputAsRows[i], " ")

		rowAsInts := []int{}
		for _, number := range stringNumbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			rowAsInts = append(rowAsInts, n)
		}

		formattedInput = append(formattedInput, rowAsInts)
	}
	return formattedInput
}

func checkAscending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] >= row[i+1] {
			return false
		}
	}
	// Check if the difference between the all numbers is between 1 and 3
	for i := 0; i < len(row)-1; i++ {
		if row[i+1]-row[i] > 3 {
			return false
		}
	}

	return true
}

func checkDescending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] <= row[i+1] {
			return false
		}
	}
	// Check if the difference between the all numbers is between 1 and 3
	for i := 0; i < len(row)-1; i++ {
		if row[i]-row[i+1] > 3 {
			return false
		}
	}

	return true
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:s]...)
	newSlice = append(newSlice, slice[s+1:]...)
	return newSlice
}
