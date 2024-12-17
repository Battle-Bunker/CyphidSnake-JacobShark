package main

import (
		"github.com/BattlesnakeOfficial/rules"
		"github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicReachableSpace returns a score based on the total number of spaces
// reachable by our team's snakes through flood fill
func HeuristicMoveSpace(snapshot agent.GameSnapshot) float64 {
		totalSpace := 0.0

		// For each snake in our team
		for _, snake := range snapshot.YourTeam() {
				if !snake.Alive() {
						continue
				}

				// Calculate reachable spaces from this snake's head
				reachable := floodFill(snapshot, snake.Head(), makeOccupiedMap(snapshot))
				totalSpace += float64(reachable)
		}

		return totalSpace
}

// floodFill returns the number of spaces reachable from the start point
func floodFill(snapshot agent.GameSnapshot, start rules.Point, occupied map[rules.Point]bool) int {
		if occupied[start] {
				return 0
		}

		count := 1
		occupied[start] = true

		// Check all four adjacent spaces (no diagonals)
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

// makeOccupiedMap creates a map of occupied spaces from snake bodies
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
