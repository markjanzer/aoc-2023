package lib

import (
	"fmt"
	"strings"
)

func StringToGrid(input string) (grid [][]byte) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	return
}

func GridToString(grid [][]byte) string {
	var lines []string
	for _, line := range grid {
		lines = append(lines, string(line))
	}

	return strings.Join(lines, "\n")
}

func PointInGrid(x, y int, grid [][]byte) bool {
	return IndexInSlice(y, grid) && IndexInSlice(x, grid[y])
}

func PrintGrid(grid [][]byte) {
	for y := range grid {
		fmt.Println(string(grid[y]))
	}
}

// Flips the x and y axis of a grid
func FlipGrid(grid [][]byte) [][]byte {
	// Dimensions of the original grid
	x := len(grid[0])
	y := len(grid)

	// Create a new grid with flipped dimensions
	newGrid := make([][]byte, x)
	for i := range newGrid {
		newGrid[i] = make([]byte, y)
	}

	// Assign values from the original grid to the new grid
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			newGrid[j][i] = grid[i][j]
		}
	}

	return newGrid
}

func ReverseGrid(grid [][]byte) [][]byte {
	for i := range grid {
		grid[i] = ReverseSlice(grid[i])
	}

	return grid
}

func GridAreEqual(grid1, grid2 [][]byte) bool {
	if len(grid1) != len(grid2) {
		return false
	}

	for i := range grid1 {
		if !EqualSlices(grid1[i], grid2[i]) {
			return false
		}
	}

	return true
}
