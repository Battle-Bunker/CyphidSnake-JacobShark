
package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

func HeuristicFloodFill(snapshot agent.GameSnapshot) float64 {
	// Get your snake
	snake := snapshot.You()
	if !snake.Alive() {
		return 0
	}

	width, height := snapshot.Width(), snapshot.Height()
	// Create visited grid
	visited := make([][]bool, width)
	for i := range visited {
		visited[i] = make([]bool, height)

	}

	head := snake.Head()
	// Mark all snake bodies as walls (including own snake and teammates)
	for _, snake := range snapshot.AllSnakes() {
		for _, part := range snake.Body() {
			// Skip marking the head position as visited
			if part != head && part.X >= 0 && part.X < width && part.Y >= 0 && part.Y < height {
				visited[part.X][part.Y] = true
			}
		}
	}

	// Get accessible space from head
	accessibleSpace := floodFill(head.X, head.Y, width, height, visited)

	return float64(accessibleSpace)
}

func floodFill(x, y, width, height int, visited [][]bool) int {
		// Check bounds and visited state
		if x < 0 || x >= width || y < 0 || y >= height || visited[x][y] {
				return 0
		}

		// Mark as visited
		visited[x][y] = true
		count := 1

		// Check adjacent cells
		count += floodFill(x+1, y, width, height, visited)
		count += floodFill(x-1, y, width, height, visited)
		count += floodFill(x, y+1, width, height, visited)
		count += floodFill(x, y-1, width, height, visited)

		return count
}