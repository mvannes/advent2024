package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	Y int
	X int
}

type Tile struct {
	Location Location
	Value    int
}

type Stack struct {
	items []Tile
}

func (s *Stack) add(loc Tile) {
	s.items = append(s.items, loc)
}

func (s *Stack) pop() Tile {
	removeIndex := len(s.items) - 1
	item := s.items[removeIndex]
	newItems := append([]Tile(nil), s.items[:removeIndex]...)
	s.items = newItems
	return item
}

func (s *Stack) empty() bool {
	return len(s.items) == 0
}

var empty interface{}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]Tile
	var trailHeads []Location
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		// create additional impassible barriers at sides of the grid
		row := []Tile{}
		for i, val := range strings.Split(line, "") {
			// impassible
			loc := Location{
				Y: lineNumber,
				X: i,
			}
			if val == "." {
				row = append(row, Tile{Value: 999, Location: loc})
				continue
			}

			intVal, err := strconv.Atoi(val)
			handleError(err)
			row = append(row, Tile{Value: intVal, Location: loc})
			if intVal == 0 {
				trailHeads = append(trailHeads, Location{
					Y: lineNumber,
					X: i,
				})
			}
		}

		grid = append(grid, row)
		lineNumber++
	}

	sumScore := 0
	for _, trailHead := range trailHeads {

		visitedList := []Tile{}
		stack := Stack{
			items: []Tile{grid[trailHead.Y][trailHead.X]},
		}
		for !stack.empty() {
			item := stack.pop()
			visitedList = append(visitedList, item)

			for _, adjacent := range AdjacentTiles(grid, item) {
				stack.add(adjacent)
			}
		}
		score := 0

		for _, tile := range visitedList {
			if tile.Value == 9 {
				score++
			}
		}
		sumScore += score
	}
	fmt.Println(sumScore)
}

func AdjacentTiles(grid [][]Tile, tile Tile) []Tile {
	result := make([]Tile, 0)

	upLoc := Location{Y: tile.Location.Y - 1, X: tile.Location.X}
	downLoc := Location{Y: tile.Location.Y + 1, X: tile.Location.X}
	leftLoc := Location{Y: tile.Location.Y, X: tile.Location.X - 1}
	rightLoc := Location{Y: tile.Location.Y, X: tile.Location.X + 1}
	adjacentLocations := []Location{upLoc, downLoc, leftLoc, rightLoc}

	// place 9, the peak, will look for place 10, but that is okay as it will not exist.
	searchVal := tile.Value + 1

	for _, loc := range adjacentLocations {
		if isLocationOutOfBounds(grid, loc) {
			continue
		}
		adjacentTile := grid[loc.Y][loc.X]
		if adjacentTile.Value == searchVal {
			result = append(result, adjacentTile)
		}
	}
	return result
}

func isLocationOutOfBounds(grid [][]Tile, loc Location) bool {
	if loc.X < 0 || loc.X >= len(grid[0]) {
		return true
	}
	if loc.Y < 0 || loc.Y >= len(grid) {
		return true
	}
	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
