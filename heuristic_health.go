package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicHealth(snapshot agent.GameSnapshot) float64 {
	totalHealth := 0
	for _, snake := range snapshot.YourTeam() {
		totalHealth += snake.Health()
	}
	// remove yourself to make your health more important
	totalHealth -= snapshot.You().Health()
	// add your self, but multiply by 3 to  make more important
	totalHealth += snapshot.You().Health() * 3
	return float64(totalHealth)
}

