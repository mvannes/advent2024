package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Location []int

func (l Location) x() int {
	return l[0]
}

func (l Location) y() int {
	return l[1]
}

func (l Location) hash() string {
	return fmt.Sprintf("%d-%d", l.x(), l.y())
}

type Tile struct {
	Location  Location
	IsBlocked bool
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var empty interface{}

type Guard struct {
	Location  Location
	Visited   map[string]interface{}
	Direction Direction
}

func (g *Guard) Look() Location {
	switch g.Direction {
	case UP:
		return Location{g.Location.x(), g.Location.y() - 1}
	case DOWN:
		return Location{g.Location.x(), g.Location.y() + 1}
	case LEFT:
		return Location{g.Location.x() - 1, g.Location.y()}
	case RIGHT:
		return Location{g.Location.x() + 1, g.Location.y()}
	}
	// Obvious err value
	return Location{-999, -999}
}

func (g *Guard) Move(loc Location) {
	// Add our previous location to the visited list, we only mark it as such when we leave.
	g.Visited[g.Location.hash()] = empty
	g.Location = loc
}

func (g *Guard) Turn() {
	var direction Direction
	switch g.Direction {
	case UP:
		direction = RIGHT
	case RIGHT:
		direction = DOWN
	case DOWN:
		direction = LEFT
	case LEFT:
		direction = UP
	}

	g.Direction = direction
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := [][]Tile{}
	guard := Guard{
		Location:  nil,
		Visited:   make(map[string]interface{}),
		Direction: UP,
	}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		tilesInLine := []Tile{}
		for i, char := range strings.Split(line, "") {
			loc := Location{i, lineNumber}
			if char == "#" {
				tilesInLine = append(tilesInLine, Tile{
					Location:  loc,
					IsBlocked: true,
				})
				continue
			} else if char == "^" {
				guard.Location = loc
			}
			tilesInLine = append(tilesInLine, Tile{
				Location:  loc,
				IsBlocked: false,
			})
		}
		lineNumber++
		grid = append(grid, tilesInLine)
	}

	gridHeight := len(grid)
	gridWidth := len(grid[0])
	// Infinite loop.
	for {
		nextLocation := guard.Look()
		if isOutOfBounds(gridHeight, gridWidth, nextLocation) {
			// Move on more time to ensure we have visited the current node, as we only
			guard.Move(nextLocation)
			break
		}
		nextTile := grid[nextLocation.y()][nextLocation.x()]
		// If our next tile is blocked, turn the guard right.
		if nextTile.IsBlocked {
			guard.Turn()
			continue
		}

		guard.Move(nextLocation)
	}

	fmt.Println(len(guard.Visited))
}

func isOutOfBounds(gridHeight, gridWidth int, location Location) bool {
	return location.x() < 0 ||
		location.x() >= gridWidth ||
		location.y() < 0 ||
		location.y() >= gridHeight
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
