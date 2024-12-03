package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)

	formattedInput := strings.ReplaceAll(input, "\r\n", "")
	formattedInput = strings.ReplaceAll(input, "\n", "")

	total1 := part1(formattedInput)
	total2 := part2(formattedInput)

	fmt.Println("Part 1 total:", total1)
	fmt.Println("Part 2 total:", total2)
}

func part1(input string) (total int) {
	validRange := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	findNumbers := regexp.MustCompile(`\d{1,3}`)

	matches := validRange.FindAllString(input, -1)
	for _, match := range matches {
		numbers := findNumbers.FindAllString(match, -1)
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		total += (num1 * num2)
	}

	return total
}

func part2(input string) (total int) {
	stringsToRemove := regexp.MustCompile(`don't\(\).*?do\(\)`)
	formattedInput := stringsToRemove.ReplaceAllString(input, "")

	validRange := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	findNumbers := regexp.MustCompile(`\d{1,3}`)

	matches := validRange.FindAllString(formattedInput, -1)
	for _, match := range matches {
		numbers := findNumbers.FindAllString(match, -1)
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		total += (num1 * num2)
	}

	return total
}
