package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rxp := regexp.MustCompile("(\\d+) +(\\d+)")

	var leftLines []int
	var rightLines []int

	for scanner.Scan() {
		line := scanner.Text()

		matches := rxp.FindStringSubmatch(line)
		left, err := strconv.Atoi(matches[1])
		handleError(err)
		leftLines = append(leftLines, left)
		right, err := strconv.Atoi(matches[2])
		handleError(err)
		rightLines = append(rightLines, right)
	}

	rightLineOccurrences := map[int]int{}
	for _, right := range rightLines {
		if _, ok := rightLineOccurrences[right]; ok {
			rightLineOccurrences[right]++
		} else {
			rightLineOccurrences[right] = 1
		}
	}

	similarityScore := 0
	for _, left := range leftLines {
		score := left * rightLineOccurrences[left]
		similarityScore += score
	}
	fmt.Println(similarityScore)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
