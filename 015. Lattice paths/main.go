/*Starting in the top left corner of a 2×2 grid, and only being able to move
to the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?*/

package main

import (
	"fmt"
)

func getgrid(size int) [][]int {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		grid[i][size-1] = 1
		grid[size-1][i] = 1
	}
	return grid
}

func main() {
	//our slice will be a bit larger then the grid
	grid := getgrid(21)
	for i := 19; i >= 0; i-- {
		for j := 19; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}

	for i := 0; i < 21; i++ {
		fmt.Println(grid[i])
	}
}
