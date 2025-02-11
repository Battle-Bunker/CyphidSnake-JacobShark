package main

import (

	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
	"math"
)

func HeuristicFoodProximity(snapshot agent.GameSnapshot) float64 {
	// Initialize biggestMD to a very large value to ensure the first distance is always smaller.
	biggestMD := math.MaxFloat64
	if snapshot.You().Health() == 100 {
		return 40
	}
	// Find the closest food to the snake's head
	for i := range snapshot.Food() {
		food := snapshot.Food()[i]
		distance := manhattanDistance(rules.Point{X: snapshot.You().Head().X, Y: snapshot.You().Head().Y}, food)
		if distance < biggestMD {
			biggestMD = distance
		}
	}

	// Calculate the score.  The score is higher when the snake is closer to food,
	// and is also influenced by the snake's health. A healthier snake might get
	// a slightly higher score.
	turnMultiplier := 1 - (0.4 * float64(snapshot.Turn()) / 150)
	if turnMultiplier < 0.3 {
		turnMultiplier = 0.3
		
	}
	healthAgressivness := 0.3 // The higher the value, the more desperate it is
	return 100.0 / (biggestMD + 1) * (1 - float64(snapshot.You().Health()) / (100*healthAgressivness) * turnMultiplier)

}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) float64 {
	return math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
}