package matrices

func islandPerimeter(grid [][]int) int {
	rows := len(grid)    // outer array length
	cols := len(grid[0]) // inner array length
	perimeter := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				perimeter += 4

				// Check if top cell is also land
				if r > 0 && grid[r-1][c] == 1 {
					perimeter -= 2
				}

				// Check if left cell is also land
				if c > 0 && grid[r][c-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}

	return perimeter
}
