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
	diskmap := strings.Split(scanner.Text(), "")

	id := 0
	index := 0
	isData := true

	var data []int
	for _, dataSizeStr := range diskmap {
		dataSize, err := strconv.Atoi(dataSizeStr)
		handleError(err)

		if isData {
			for i := index; i < index+dataSize; i++ {
				data = append(data, id)
			}
			id++
		} else {
			for i := index; i < index+dataSize; i++ {
				data = append(data, -10)
			}
		}

		// Flip around, handling empty blocks
		isData = !isData
		index += dataSize
	}

	fmt.Println(data)

	dataIndex := len(data)
	firstEmptyIndex := 0
	for i := len(data) - 1; i > 0; i-- {
		// if it is empty, we can continue
		if data[i] == -10 {
			continue
		}
		dataIndex = i
		// We have to find what to swap.
		foundEmpty := false
		// Everything behind the data index must be empty.
		for j := firstEmptyIndex; j < dataIndex; j++ {
			// This time, if it is empty, set new firstEmpty.
			if data[j] == -10 {
				firstEmptyIndex = j
				foundEmpty = true
				break
			}
		}
		// We must be done, no empty index found at all.
		if !foundEmpty {
			break
		}
		// Swap them around
		data[firstEmptyIndex] = data[dataIndex]
		data[dataIndex] = -10
	}

	fmt.Println(data)
	sum := 0
	for i, val := range data {
		if val == -10 {
			continue
		}

		sum += (i * val)
	}
	fmt.Println(sum)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
