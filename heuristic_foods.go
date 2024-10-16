package main

import (
	"math"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicFoodProximity calculates a score based on the proximity of team snakes to food
func HeuristicFoodProximity(snapshot agent.GameSnapshot) float64 {
	var totalScore float64

	for _, snake := range snapshot.YourTeam() {
		if !snake.Alive() {
			continue
		}

		closestFoodDistance := closestFoodDistance(snake.Head(), snapshot.Food())
		if closestFoodDistance == math.MaxFloat64 {
			continue // No food on the board
		}

		// Score is inversely proportional to the distance to the closest food
		// We add 1 to avoid division by zero and to give some value even when on food
		snakeScore := 100.0 / (closestFoodDistance + 1)
		totalScore += snakeScore
	}

	return totalScore
}

// closestFoodDistance calculates the Manhattan distance to the closest food
func closestFoodDistance(head rules.Point, food []rules.Point) float64 {
	if len(food) == 0 {
		return math.MaxFloat64
	}

	closestDistance := math.MaxFloat64
	for _, f := range food {
		distance := manhattanDistance(head, f)
		if distance < closestDistance {
			closestDistance = distance
		}
	}
	return closestDistance
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) float64 {
	return math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
}