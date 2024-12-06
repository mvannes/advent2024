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
	return fmt.Sprintf("%d-%d", l.y(), l.x())
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
	Location Location
	// Could create a custom type that does not expose that this is a set, but can't be bothered.
	Visited   map[string]Location
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
	// Include the direction, because we it will not be a duplicate if it is not involved.
	g.Visited[fmt.Sprintf("%s-%d", loc.hash(), g.Direction)] = g.Location
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

func (g *Guard) HasVisited(location Location) bool {
	_, hasVisited := g.Visited[fmt.Sprintf("%s-%d", location.hash(), g.Direction)]
	return hasVisited
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
		Visited:   make(map[string]Location),
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

	// Even in a struct, slices and maps in GO are pass by reference.
	// So, make a copy for every time we want to iterate through.
	startGuard := Guard{
		Location:  guard.Location,
		Visited:   map[string]Location{fmt.Sprintf("%s-%d", guard.Location.hash(), UP): guard.Location},
		Direction: UP,
	}
	willLoop(grid, startGuard)
	// We've used the startGuard to find the initial path our guard will take.
	// Now, block each step in their path (because anything else will not affect their pathing anyway)
	loopingCounter := 0
	locationMap := map[string]interface{}{}
	for _, location := range startGuard.Visited {
		if _, ok := locationMap[location.hash()]; ok {
			continue
		}
		locationMap[location.hash()] = empty

		// Deep copy the grid, for similar slice reference logic
		gridCopy := copyGrid(grid)
		// block the current location
		gridCopy[location.y()][location.x()].IsBlocked = true
		attemptGuard := Guard{
			Location:  guard.Location,
			Visited:   map[string]Location{fmt.Sprintf("%s-%d", guard.Location.hash(), UP): guard.Location},
			Direction: UP,
		}

		if willLoop(gridCopy, attemptGuard) {
			loopingCounter++
		}
	}

	fmt.Println(loopingCounter)
}

func copyGrid(grid [][]Tile) [][]Tile {
	gridCopy := make([][]Tile, len(grid))
	for i, tiles := range grid {
		gridCopy[i] = make([]Tile, len(tiles))
		for j, tile := range tiles {
			gridCopy[i][j] = Tile{
				Location:  Location{tile.Location.x(), tile.Location.y()},
				IsBlocked: tile.IsBlocked,
			}
		}
	}
	return gridCopy
}

func willLoop(grid [][]Tile, guard Guard) bool {
	gridHeight := len(grid)
	gridWidth := len(grid[0])
	// Infinite loop.
	for {
		nextLocation := guard.Look()
		if isOutOfBounds(gridHeight, gridWidth, nextLocation) {
			guard.Move(nextLocation)
			return false
		}

		if guard.HasVisited(nextLocation) {
			return true
		}
		nextTile := grid[nextLocation.y()][nextLocation.x()]
		// If our next tile is blocked, turn the guard right.
		if nextTile.IsBlocked {
			guard.Turn()
			continue
		}

		guard.Move(nextLocation)
	}
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
