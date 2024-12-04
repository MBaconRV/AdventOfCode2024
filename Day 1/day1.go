package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)
	inputAsRows := strings.Split(input, "\n")

	leftColumn, rightColumn := splitAndSortInput(inputAsRows)

	totalDistance := findTotalDistance(leftColumn, rightColumn)
	fmt.Println("total distance:", totalDistance)

	similarityScore := findSimilarityScore(leftColumn, rightColumn)
	fmt.Println("similarity score:", similarityScore)
}

func splitAndSortInput(inputAsRows []string) (leftColumn, rightColumn []int) {
	for _, row := range inputAsRows {
		numbers := strings.Split(row, "   ")
		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		leftColumn = append(leftColumn, number1)
		rightColumn = append(rightColumn, number2)
	}

	slices.Sort(leftColumn)
	slices.Sort(rightColumn)

	return leftColumn, rightColumn
}

func findTotalDistance(leftColumn, rightColumn []int) (totalDistance int) {
	for i, number1 := range leftColumn {
		number2 := rightColumn[i]
		distance := number1 - number2
		if distance < 0 {
			distance = distance * -1
		}
		totalDistance += distance
	}

	return
}

func findSimilarityScore(leftColumn, rightColumn []int) (similarityScore int) {
	for _, leftNumber := range leftColumn {
		frequency := 0
		for _, rightNumber := range rightColumn {
			if leftNumber == rightNumber {
				frequency++
			}
		}
		if frequency > 0 {
			similarityScore += (frequency * leftNumber)
		}
	}

	return
}
