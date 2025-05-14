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
		distance := ManhattanDistance(snapshot.You().Head(), food)
		if float64(distance) < biggestMD {
			biggestMD = float64(distance)
		}
	}

	// Calculate the score.  The score is higher when the snake is closer to food,
	// and is also influenced by the snake's health. A healthier snake might get
	// a slightly higher score.
	// turnMultiplier := 1 - (0.4 * float64(snapshot.Turn()) / 150)

	// Assuming that snapshot.AllSnakes() returns a slice of all snakes and 
	// each snake has an IsAlive() method to check if a snake is alive.
	var aliveSnakes float64 = 0
	for _, snake := range snapshot.AllSnakes() {
		if snake.Alive() { // Check if the snake is alive.
			aliveSnakes++
		}
	}
	
	snakeMultiplier := 1.0
	totalSnakes := float64(len(snapshot.AllSnakes()))
	if totalSnakes > 0 {
		snakeFraction := float64(aliveSnakes/totalSnakes)
		if snakeFraction == 1 { // everyone is in
			snakeMultiplier = 0.5
		} else if snakeFraction >= 5/6 {
			snakeMultiplier = 0.4
		} else if snakeFraction >= 4/6 {
			snakeMultiplier = 0.3
		} else if aliveSnakes >= 3/6 { // half the snakes are dead
			snakeMultiplier = 0.2
		} else if aliveSnakes >= 2/6 {
			snakeMultiplier = 0.2
			
		}
	}
	
	
	healthAgressivness := 0.3 // The higher the value, the more desperate it is
	return ((100.0 / (biggestMD + 1) * (1 - float64(snapshot.You().Health()) / (100*healthAgressivness) * snakeMultiplier)) / 11) * float64(snapshot.Height())

}




// ManhattanDistance calculates the Manhattan distance between two points.
func ManhattanDistance(point1, point2 rules.Point) int {
    return int(math.Abs(float64(point1.X-point2.X)) + math.Abs(float64(point1.Y-point2.Y)))
}

