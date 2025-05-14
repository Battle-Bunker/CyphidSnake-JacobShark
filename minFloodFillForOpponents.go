package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func MinFLoodfillForEnemy(snapshot agent.GameSnapshot) float64 {
    MyHead := snapshot.You().Head()
    // Find closest enemy
    minDist := 1000000000
    minSnake := snapshot.Opponents()[0]
    for _, snake := range snapshot.Opponents() {
        if ManhattanDistance(MyHead, snake.Head()) < minDist {
            minDist = ManhattanDistance(MyHead, snake.Head())
            minSnake = snake
        }
    }

    
    return float64(floodFill(minSnake.Head().X, minSnake.Head().Y, snapshot.Width(), snapshot.Height(), make([][]bool, snapshot.Width()))*-1)
    
}