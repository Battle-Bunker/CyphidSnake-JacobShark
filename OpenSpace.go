
package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicOpenSpaceFloodFill returns a score based on the average open space available
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
	directions := []rules.SnakeMove{
		{Move: "up"},
		{Move: "down"},
		{Move: "left"},
		{Move: "right"},
	}

	head := snake.Head()
	for _, dir := range directions {
		var nextHead rules.Point
		switch dir.Move {
		case "up":
			nextHead = rules.Point{X: head.X, Y: head.Y + 1}
		case "down":
			nextHead = rules.Point{X: head.X, Y: head.Y - 1}
		case "left":
			nextHead = rules.Point{X: head.X - 1, Y: head.Y}
		case "right":
			nextHead = rules.Point{X: head.X + 1, Y: head.Y}
		}

		// Check if the point is within bounds and empty
		if nextHead.X >= 0 && nextHead.X < snapshot.Width() &&
			nextHead.Y >= 0 && nextHead.Y < snapshot.Height() {
			isOccupied := false
			for _, otherSnake := range snapshot.AllSnakes() {
				for _, segment := range otherSnake.Body() {
					if segment.X == nextHead.X && segment.Y == nextHead.Y {
						isOccupied = true
						break
					}
				}
				if isOccupied {
					break
				}
			}
			if !isOccupied {
				openSpace++
			}
		}
	}
	return openSpace
}
