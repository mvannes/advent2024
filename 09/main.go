package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Id         int
	StartIndex int
	EndIndex   int
}

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
	var files []File
	for _, dataSizeStr := range diskmap {
		dataSize, err := strconv.Atoi(dataSizeStr)
		handleError(err)

		if isData {
			for i := index; i < index+dataSize; i++ {
				data = append(data, id)
			}
			files = append(files, File{
				Id:         id,
				StartIndex: index,
				EndIndex:   index + dataSize,
			})
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

	// Go backwards through the files
	for i := len(files) - 1; i > 0; i-- {
		firstEmptyIndex := 0
		file := files[i]
		spaceRequired := file.EndIndex - file.StartIndex
		// We have to find what to swap.
		foundEmptyLargeEnough := false
		for j := firstEmptyIndex; j < len(data)-1; j++ {
			if data[j] != -10 {
				continue
			}
			firstEmptyIndex = j
			spaceFound := true
			// Look ahead for as much as we need, stopping if we find non empty space (it won't fit.
			for s := 0; s < spaceRequired; s++ {
				possibleIndex := firstEmptyIndex + s
				if possibleIndex >= len(data)-1 || data[possibleIndex] != -10 {
					// Set to skip the bits we checked, as those must all not fit, since the smallest
					// does not ift.
					j = firstEmptyIndex + s
					spaceFound = false
					break
				}
			}

			// Break out if we found space.
			if spaceFound {
				foundEmptyLargeEnough = true
				break
			}
		}
		// could not find empty space that is not behind our file.
		if !foundEmptyLargeEnough || firstEmptyIndex >= file.StartIndex {
			continue
		}
		// Swap them around
		for s := 0; s < spaceRequired; s++ {
			data[firstEmptyIndex+s] = data[file.StartIndex+s]
			data[file.StartIndex+s] = -10
		}
	}

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
