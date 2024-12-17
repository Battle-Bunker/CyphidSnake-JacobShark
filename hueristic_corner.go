package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicCorner(snapshot agent.GameSnapshot) float64 {
    curPos := [2]int{snapshot.You().Head().X, snapshot.You().Head().Y}
    badPos := [][2]int{{0, 0}, {10, 0}, {0, 10}, {10, 10}}

    for _, ar := range badPos {
        if curPos[0] == ar[0] && curPos[1] == ar[1] {
            return -10
        }
    }

    return 0
}
