package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by Jacob)
func HeuristicSelfDeath(snapshot agent.GameSnapshot) float64 {
  return float64(snapshot.You().Health())-float64(50)
}
