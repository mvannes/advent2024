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

	rxp := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")

	var multiples [][]int
	for scanner.Scan() {
		line := scanner.Text()
		matches := rxp.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			left, err := strconv.Atoi(match[1])
			handleError(err)
			right, err := strconv.Atoi(match[2])
			handleError(err)
			multiples = append(multiples, []int{left, right})
		}
	}

	var sum int

	for _, pair := range multiples {
		sum += pair[0] * pair[1]
	}

	fmt.Println(sum)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
