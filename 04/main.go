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
			// we start parsing on finding an A
			if parsedChar != "A" {
				continue
			}
			if isXmas(lines, []int{i, j}) {
				xmasCounter++
			}
		}
	}

	fmt.Println(xmasCounter)
}

func isXmas(lines [][]string, aLocation []int) bool {
	y := aLocation[0]
	x := aLocation[1]

	a := lines[y][x]
	leftUp := lines[y-1][x-1]
	rightUp := lines[y-1][x+1]
	leftDown := lines[y+1][x-1]
	rightDown := lines[y+1][x+1]

	diagonalTopToBot := leftUp + a + rightDown
	diagonalBotToTop := leftDown + a + rightUp

	hasMasTopToBot := diagonalTopToBot == "MAS" || diagonalTopToBot == "SAM"
	hasMasBotToTop := diagonalBotToTop == "MAS" || diagonalBotToTop == "SAM"

	return hasMasBotToTop && hasMasTopToBot
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
