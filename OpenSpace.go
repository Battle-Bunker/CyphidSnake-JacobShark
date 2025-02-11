
package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

func HeuristicFloodFill(snapshot agent.GameSnapshot) float64 {
	// Get your snake
	snake := snapshot.You()
	if !snake.Alive() {
		return 0
	}

	// Create visited map
	visited := make(map[rules.Point]bool)

	head := snake.Head()
	// Mark all snake bodies as walls (including own snake and teammates)
	for _, snake := range snapshot.AllSnakes() {
		for _, part := range snake.Body() {
			// Skip marking the head position as visited
			if part != head {
				visited[part] = true
			}
		}
	}

	// Get accessible space from head
	accessibleSpace := floodFill(snapshot, head, visited)

	return float64(accessibleSpace)
}

func floodFill(snapshot agent.GameSnapshot, point rules.Point, visited map[rules.Point]bool) int {
	// Check bounds
	if point.X < 0 || point.X >= snapshot.Width() || point.Y < 0 || point.Y >= snapshot.Height() {
		return 0
	}

	// Check if already visited
	if visited[point] {
		return 0
	}

	// Mark as visited
	visited[point] = true
	count := 1

	// Check adjacent cells
	directions := []rules.Point{
		{X: point.X + 1, Y: point.Y},
		{X: point.X - 1, Y: point.Y},
		{X: point.X, Y: point.Y + 1},
		{X: point.X, Y: point.Y - 1},
	}

	for _, dir := range directions {
		count += floodFill(snapshot, dir, visited)
	}

	return count
}
