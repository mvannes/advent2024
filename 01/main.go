package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

	sort.Ints(leftLines)
	sort.Ints(rightLines)

	sum := 0
	for i := 0; i < len(leftLines); i++ {
		left := leftLines[i]
		right := rightLines[i]

		difference := int(math.Abs(float64(left - right)))

		fmt.Println(left, right, difference)
		sum += difference
	}
	fmt.Println(sum)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
