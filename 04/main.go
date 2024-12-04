package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		// prepend and append dots to make searching for xmas easier.
		// Means we can start at index 4, and not concern ourselves with the bounds of the puzzle input.
		line = "...." + line + "...."

		lines = append(lines, strings.Split(line, ""))
	}

	additionalLine := strings.Split(strings.Repeat(".", len(lines[0])), "")
	topBotPadding := [][]string{additionalLine, additionalLine, additionalLine, additionalLine}

	lines = append(topBotPadding, lines...)
	lines = append(lines, topBotPadding...)

	xmasCounter := 0

	// Start at 4, because we have padded the front and back with empty liens
	for i := 4; i < len(lines)-4; i++ {
		// same padding offset here.
		for j := 4; j < len(lines[i])-4; j++ {
			parsedChar := lines[i][j]
			// we start parsing on finding an X.
			if parsedChar != "X" {
				continue
			}
			xmasCounter += xmasCount(lines, []int{i, j})
		}
	}

	fmt.Println(xmasCounter)
}

func xmasCount(lines [][]string, xLocation []int) int {

	y := xLocation[0]
	x := xLocation[1]

	forward := ""
	backward := ""
	up := ""
	upLeft := ""
	upRight := ""
	down := ""
	downLeft := ""
	downRight := ""
	for i := 0; i < 4; i++ {
		forward += lines[y][x+i]
		backward += lines[y][x-i]
		up += lines[y-i][x]
		upLeft += lines[y-i][x-i]
		upRight += lines[y-i][x+i]
		down += lines[y+i][x]
		downLeft += lines[y+i][x-i]
		downRight += lines[y+i][x+i]
	}

	counter := 0
	if forward == "XMAS" {
		counter++
	}

	if backward == "XMAS" {
		counter++
	}
	if up == "XMAS" {
		counter++
	}
	if upLeft == "XMAS" {
		counter++
	}
	if upRight == "XMAS" {
		counter++
	}
	if down == "XMAS" {
		counter++
	}
	if downLeft == "XMAS" {
		counter++
	}
	if downRight == "XMAS" {
		counter++
	}

	return counter
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
