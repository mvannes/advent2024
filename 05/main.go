package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// map of BEFORE -> What it needs to be before.
	pageOrderings := map[int][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Done reading page orderings.
			break
		}
		split := strings.Split(line, "|")
		before, err := strconv.Atoi(split[0])
		handleError(err)
		after, err := strconv.Atoi(split[1])
		handleError(err)

		if afterSlice, ok := pageOrderings[before]; ok {
			afterSlice = append(afterSlice, after)
			pageOrderings[before] = afterSlice
		} else {
			pageOrderings[before] = []int{after}
		}
	}

	var pageNumberUpdates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var pageNumbers []int
		pageNumberStrings := strings.Split(line, ",")
		for _, pageNumberStr := range pageNumberStrings {
			pageNumber, err := strconv.Atoi(pageNumberStr)
			handleError(err)
			pageNumbers = append(pageNumbers, pageNumber)
		}

		pageNumberUpdates = append(pageNumberUpdates, pageNumbers)
	}

	middleNumberSum := 0
	for _, update := range pageNumberUpdates {
		sorted := sort.SliceIsSorted(update, func(i, j int) bool {
			a := update[i]
			b := update[j]

			if afters, ok := pageOrderings[a]; ok {
				for _, after := range afters {
					if b == after {
						return true
					}
				}
			}
			if afters, ok := pageOrderings[b]; ok {
				for _, after := range afters {
					if a == after {
						return false
					}
				}
			}
			return false
		})

		if sorted {
			fmt.Println(update, update[len(update)/2])
			middleNumberSum += update[len(update)/2]
		}
	}
	fmt.Println(middleNumberSum)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
