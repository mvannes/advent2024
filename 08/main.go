package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Location struct {
	x int
	y int
}

type Antenna struct {
	Location Location
	Signal   string
}

var empty interface{}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]string
	// Immediately group by type
	antennasBySignal := map[string][]Antenna{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)

		for i, val := range row {
			if val == "." {
				continue
			}

			antenna := Antenna{
				Location: Location{
					y: lineNumber,
					x: i,
				},
				Signal: val,
			}
			if curr, ok := antennasBySignal[val]; ok {
				antennasBySignal[val] = append(curr, antenna)
			} else {
				antennasBySignal[val] = []Antenna{antenna}
			}
		}
		lineNumber++
	}

	height := len(grid)
	width := len(grid[0])

	nodeLocations := []Location{}

	for _, antennas := range antennasBySignal {
		for i := 0; i < len(antennas); i++ {
			a := antennas[i]
			// Compare the antenna with every other antenna of its kind.
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}
				b := antennas[j]

				var xDiff int
				var goRight bool
				// a is more right than b
				if a.Location.x > b.Location.x {
					xDiff = a.Location.x - b.Location.x
					goRight = true
				} else {
					xDiff = b.Location.x - a.Location.x
					goRight = false
				}

				var yDiff int
				var goDown bool
				// a is more Up than b
				if a.Location.y > b.Location.y {
					yDiff = a.Location.y - b.Location.y
					goDown = true
				} else {
					yDiff = b.Location.y - a.Location.y
					goDown = false
				}

				location := a.Location

				nodeLocations = append(nodeLocations, location)
				for {
					var antiX int
					if goRight {
						antiX = location.x + xDiff
					} else {
						antiX = location.x - xDiff
					}

					var antiY int
					if goDown {
						antiY = location.y + yDiff
					} else {
						antiY = location.y - yDiff
					}

					location = Location{
						x: antiX,
						y: antiY,
					}

					if isOutOfBounds(height, width, location) {
						break
					}
					nodeLocations = append(nodeLocations, location)
				}
			}
		}
	}

	uniqueLocations := 0
	for _, location := range nodeLocations {
		if isOutOfBounds(height, width, location) {
			continue
		}
		if grid[location.y][location.x] == "#" {
			continue
		}
		grid[location.y][location.x] = "#"
		uniqueLocations++
	}

	fmt.Println(uniqueLocations)
	//for _, row := range grid {
	//	fmt.Println(strings.Join(row, ""))
	//}
}

func isOutOfBounds(height, width int, location Location) bool {
	if location.y < 0 || location.y >= height {
		return true
	}
	if location.x < 0 || location.x >= width {
		return true
	}
	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
