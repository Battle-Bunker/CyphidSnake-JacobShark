package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	_ "github.com/BattlesnakeOfficial/rules"
)

// HeuristicOpenSpace returns a score based on the average open space available
// to each allied snake. For each snake, it looks at the immediate squares
// available for forward movement and assigns a higher score when more squares
// are accessible.
func HeuristicOpenSpaceFloodFill(snapshot agent.GameSnapshot) float64 {

	var totalOpenSpace float64 = 0.0
	for _, snake := range snapshot.Snakes() {
		openSpace := countOpenSpace(snake, snapshot)
		totalOpenSpace += float64(openSpace)
	}
	return totalOpenSpace / float64(len(snapshot.Snakes()))
}

func countOpenSpace(snake agent.SnakeSnapshot, snapshot agent.GameSnapshot) int {
	openSpace := 0
	for _, direction := range []string{"up", "down", "left", "right"} {
		nextHead := snake.Head.Move(direction)
		if nextHead.IsPassable() && snapshot.Board()[nextHead.Y][nextHead.X] == agent.Empty {
			openSpace++
		}
	}
	return openSpace
}
