package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	Anwser int
	Values []int
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var calibrations []Calibration
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")
		answer, err := strconv.Atoi(split[0])
		handleError(err)
		var values []int
		for _, valueStr := range strings.Split(split[1], " ") {
			value, err := strconv.Atoi(valueStr)
			handleError(err)
			values = append(values, value)
		}

		calibrations = append(calibrations, Calibration{
			Anwser: answer,
			Values: values,
		})
	}

	validTotal := 0
	for _, calibration := range calibrations {
		amountOfOperatorPositions := len(calibration.Values) - 1
		combinations := [][]string{{}}

		for i := 0; i < amountOfOperatorPositions; i++ {
			// For every position, loop over the previous one.
			newCombinations := [][]string{}
			for _, c := range combinations {
				add := append([]string(nil), c...)
				add = append(add, "+")

				mul := append([]string(nil), c...)
				mul = append(mul, "*")

				concat := append([]string(nil), c...)
				concat = append(concat, "||")
				newCombinations = append(newCombinations, add, mul, concat)
			}
			combinations = newCombinations
		}
		// Check all combinations
		foundValid := false
		for _, combination := range combinations {
			if isValidCombination(calibration, combination) {
				foundValid = true
				break
			}
		}
		if foundValid {
			validTotal += calibration.Anwser
		}
	}

	fmt.Println(validTotal)
}

func isValidCombination(calibration Calibration, combination []string) bool {
	total := calibration.Values[0]

	for i, operator := range combination {
		switch operator {
		case "*":
			// Plus one, as we look towards the next one / we loop over the operators, not the numbers
			total *= calibration.Values[i+1]
		case "+":
			total += calibration.Values[i+1]
		case "||":
			totalAsStr := fmt.Sprintf("%d%d", total, calibration.Values[i+1])
			newTotal, err := strconv.Atoi(totalAsStr)
			handleError(err)
			total = newTotal
		}

		// Early escape out of branches that cross over the boundary.
		if total > calibration.Anwser {
			return false
		}
	}
	return total == calibration.Anwser
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
