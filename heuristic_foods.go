package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)
import (
	"math"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicFood(gs agent.GameSnapshot) float64 {
	you := gs.You()
	head := you.Head()
	food := gs.Food()

	if len(food) == 0 {
		return math.Inf(1) // Return positive infinity if there's no food
	}

	lowestDist := math.Inf(1)

	for _, foodPoint := range food {
		dist := euclideanDistance(head, foodPoint)
		if dist < lowestDist {
			lowestDist = dist
		}
	}

	return lowestDist

}
func manhattanDistance(p1 rules.Point, p2 rules.Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

func euclideanDistance(p1, p2 rules.Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
