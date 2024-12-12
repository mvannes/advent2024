package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Location struct {
	Y int
	X int
}

type Tile struct {
	Location Location
	Value    string
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

type Plot struct {
	tiles []PlotTile
}

type PlotTile struct {
	Tile                  Tile
	NonPlotAdjacencyCount int // amount of non this plot tiles surrounding this one.
}

var empty interface{}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]Tile
	lineNumber := 0
	var locs []Location
	for scanner.Scan() {
		line := scanner.Text()
		// create additional impassible barriers at sides of the grid
		row := []Tile{}
		for i, val := range strings.Split(line, "") {
			loc := Location{
				Y: lineNumber,
				X: i,
			}
			row = append(row, Tile{Value: val, Location: loc})
			locs = append(locs, loc)
		}

		grid = append(grid, row)
		lineNumber++
	}

	fullVisited := map[Location]interface{}{}
	isBuildingPlot := false
	var plot Plot
	var plots []Plot

	for _, loc := range locs {
		// Skip indices that we've already explored.
		// Since we do a DFS to get each full plot, I'm not sure how else to
		// find the next plot without just attempting.
		if _, ok := fullVisited[loc]; ok {
			continue
		}

		if !isBuildingPlot {
			plot = Plot{tiles: make([]PlotTile, 0)}
			isBuildingPlot = true
		}

		visitedList := []Tile{}
		stack := Stack{
			items: []Tile{grid[loc.Y][loc.X]},
		}
		for !stack.empty() {
			item := stack.pop()
			if _, ok := fullVisited[item.Location]; ok {
				continue
			}
			fullVisited[item.Location] = empty

			visitedList = append(visitedList, item)
			adjacencies, amountOfInvalidAdjacencies := AdjacentTiles(grid, item)
			plot.tiles = append(plot.tiles, PlotTile{
				Tile:                  item,
				NonPlotAdjacencyCount: amountOfInvalidAdjacencies,
			})
			for _, adjacent := range adjacencies {
				stack.add(adjacent)
			}
		}
		isBuildingPlot = false
		plots = append(plots, plot)
	}

	counter := 0
	for _, p := range plots {
		area := len(p.tiles)
		border := 0
		for _, tile := range p.tiles {
			border += tile.NonPlotAdjacencyCount
		}

		counter += (area * border)
	}

	fmt.Println(counter)
}

func AdjacentTiles(grid [][]Tile, tile Tile) ([]Tile, int) {
	result := make([]Tile, 0)

	upLoc := Location{Y: tile.Location.Y - 1, X: tile.Location.X}
	downLoc := Location{Y: tile.Location.Y + 1, X: tile.Location.X}
	leftLoc := Location{Y: tile.Location.Y, X: tile.Location.X - 1}
	rightLoc := Location{Y: tile.Location.Y, X: tile.Location.X + 1}
	adjacentLocations := []Location{upLoc, downLoc, leftLoc, rightLoc}

	invalidPositions := 0
	for _, loc := range adjacentLocations {
		if isLocationOutOfBounds(grid, loc) {
			invalidPositions++
			continue
		}
		adjacentTile := grid[loc.Y][loc.X]
		// Only take if they are the same plant
		if adjacentTile.Value == tile.Value {
			result = append(result, adjacentTile)
		} else {
			invalidPositions++
		}
	}
	return result, invalidPositions
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
