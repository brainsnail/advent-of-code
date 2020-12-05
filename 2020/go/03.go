package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputFile = "../inputs/03.txt"

func main() {

	var grid []string

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	var treeHits = sled(grid, MovementRules{3, 1})
	fmt.Printf("You hit this many trees: %d \n", treeHits)
}

type MovementRules struct {
	X int
	Y int
}

// Repeat x,y position movements across a grid til you hit the floor of the Y axis
func sled(grid []string, direction MovementRules) int {

	// Keep a running total of times you hit a 'tree'
	var treeHits int

	// Keep track of current position in the grid
	var pos MovementRules

	size := MovementRules{len(grid[0]), len(grid)}

	// Keeps going to the right til you hit the bottom of the 'slope'
	for pos.Y < size.Y {
		if grid[pos.Y][pos.X%size.X] == '#' {
			treeHits++ // Ouch
		}
		pos.X += direction.X
		pos.Y += direction.Y
	}
	return treeHits
}
