package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	stoneStrings := strings.Split(line, " ")
	// Amount of stones for each stone string.
	stoneMap := map[string]int{}
	// Init the stone map with initial input
	for _, stoneStr := range stoneStrings {
		if val, ok := stoneMap[stoneStr]; ok {
			stoneMap[stoneStr] = val + 1
		} else {
			stoneMap[stoneStr] = 1
		}
	}

	// Per stone, cache what stones it would result in when blinking.
	stoneBlinkResultCache := map[string][]string{}

	amountOfBlinks := 0
	for amountOfBlinks < 25 {
		newStoneMap := map[string]int{}
		for stoneStr, amountOfStones := range stoneMap {
			var newStones []string

			if cachedStones, ok := stoneBlinkResultCache[stoneStr]; ok {
				// If we have the result, do not calculate.
				newStones = append([]string(nil), cachedStones...)
			} else {
				newStones = calculateStones(stoneStr)
				// Store result in the cache.
				stoneBlinkResultCache[stoneStr] = append([]string(nil), newStones...)
			}

			for _, newStone := range newStones {
				if val, ok := newStoneMap[newStone]; ok {
					newStoneMap[newStone] = val + amountOfStones
				} else {
					newStoneMap[newStone] = amountOfStones
				}
			}
		}

		// Override the old stonemap and increment blinks
		stoneMap = newStoneMap
		amountOfBlinks++
	}

	stoneCounter := 0
	for _, stoneAmount := range stoneMap {
		stoneCounter += stoneAmount
	}
	fmt.Println(stoneCounter)
}

func calculateStones(stoneStr string) []string {
	// if odd amount of characters, do times 2024 if not 0
	if len(stoneStr)%2 == 1 {
		stone, err := strconv.Atoi(stoneStr)
		handleError(err)
		if stone == 0 {
			return []string{"1"}
		} else {
			return []string{strconv.Itoa(stone * 2024)}
		}
	}

	stoneHalf := len(stoneStr) / 2
	leftSide := stoneStr[:stoneHalf]
	rightSide := stoneStr[stoneHalf:]

	return []string{truncateZeros(leftSide), truncateZeros(rightSide)}
}

func truncateZeros(str string) string {
	withoutZeros := strings.TrimLeft(str, "0")
	// if it was only zeros, add the zero back in
	if withoutZeros == "" {
		withoutZeros = "0"
	}
	return withoutZeros
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
