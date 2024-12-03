package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	rxp := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)|(do\\(\\))|(don't\\(\\))")

	var multiples [][]int

	doing := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := rxp.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			switch {
			case strings.HasPrefix(match[0], "don't"):
				doing = false
				continue
			case strings.HasPrefix(match[0], "do"):
				doing = true
				continue
			}
			if !doing {
				continue
			}
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
