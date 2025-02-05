
package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicOpenSpaceFloodFill returns a score based on the average open space available
// to each allied snake using flood fill algorithm for more accurate space calculation
func HeuristicOpenSpaceFloodFill(snapshot agent.GameSnapshot) float64 {
	var totalOpenSpace float64 = 0.0
	
	// Only consider your team's snakes
	for _, snake := range snapshot.YourTeam() {
		if !snake.Alive() {
			continue
		}
		// Use flood fill to calculate reachable spaces
		openSpace := floodFill(snapshot, snake.Head(), makeOccupiedMap(snapshot))
		totalOpenSpace += float64(openSpace)
	}
	
	numSnakes := len(snapshot.YourTeam())
	if numSnakes == 0 {
		return 0
	}
	return totalOpenSpace / float64(numSnakes)
}

func floodFill(snapshot agent.GameSnapshot, start rules.Point, occupied map[rules.Point]bool) int {
	if occupied[start] {
		return 0
	}

	count := 1
	occupied[start] = true

	// Check all four adjacent spaces
	directions := []rules.Point{
		{X: 0, Y: 1},  // up
		{X: 0, Y: -1}, // down
		{X: 1, Y: 0},  // right
		{X: -1, Y: 0}, // left
	}

	for _, dir := range directions {
		next := rules.Point{
			X: start.X + dir.X,
			Y: start.Y + dir.Y,
		}

		// Check if the point is within bounds
		if next.X >= 0 && next.X < snapshot.Width() &&
			next.Y >= 0 && next.Y < snapshot.Height() &&
			!occupied[next] {
			count += floodFill(snapshot, next, occupied)
		}
	}

	return count
}

func makeOccupiedMap(snapshot agent.GameSnapshot) map[rules.Point]bool {
	occupied := make(map[rules.Point]bool)

	// Mark all snake body segments as occupied
	for _, snake := range snapshot.AllSnakes() {
		if !snake.Alive() {
			continue
		}
		for _, segment := range snake.Body() {
			occupied[segment] = true
		}
	}

	// Mark hazards as occupied
	for _, hazard := range snapshot.Hazards() {
		occupied[hazard] = true
	}

	return occupied
}
