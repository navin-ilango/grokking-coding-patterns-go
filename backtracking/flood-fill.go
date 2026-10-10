package backtracking

// floodFill serves as the entry point and initial check.
func floodFill(grid [][]int, sr, sc, target int) [][]int {
	// Guard clause: If the target pixel already has the new color,
	// changing it would cause infinite recursion (stack overflow).
	startColor := grid[sr][sc]
	if startColor == target {
		return grid
	}

	// Capture grid dimensions
	rows := len(grid)
	cols := len(grid[0])

	// Kick off the recursive helper
	dfs(grid, sr, sc, startColor, target, rows, cols)

	return grid
}

// dfs handles the recursive traversal.
func dfs(grid [][]int, r, c, startColor, target, rows, cols int) {
	// 1. Base Case: Check out-of-bounds bounds to prevent runtime panics.
	if r < 0 || r >= rows || c < 0 || c >= cols {
		return
	}

	// 2. Base Case: Only modify pixels matching the original color.
	if grid[r][c] != startColor {
		return
	}

	// 3. Process: Update the color of the current pixel
	grid[r][c] = target

	// 4. Recurse: Explore 4-directional neighbors
	dfs(grid, r+1, c, startColor, target, rows, cols) // Down
	dfs(grid, r-1, c, startColor, target, rows, cols) // Up
	dfs(grid, r, c+1, startColor, target, rows, cols) // Right
	dfs(grid, r, c-1, startColor, target, rows, cols) // Left
}
