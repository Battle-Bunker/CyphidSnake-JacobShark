package main

import (
    _ "github.com/BattlesnakeOfficial/rules"
    _ "github.com/Battle-Bunker/cyphid-snake/agent"
    "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicOpenSpace returns a score based on the average open space available
// to each allied snake. For each snake, it looks at the immediate squares
// available for forward movement and assigns a higher score when more squares
// are accessible.
func HeuristicOpenSpace(snapshot agent.GameSnapshot) float64 {
    totalScore := 0.0
    allySnakes := snapshot.YourTeam()

    // If no allied snakes are alive, return minimum score
    if len(allySnakes) == 0 {
        return 0.0
    }

    for _, snake := range allySnakes {
        if !snake.Alive() {
            continue
        }

        // Count the number of forward moves available
        moves := snake.ForwardMoves()
        moveCount := len(moves)

        // Score based on available moves:
        // 1 move = 1 point (narrow corridor)
        // 2 moves = 4 points (corner or wider corridor)
        // 3 moves = 9 points (open space with one blocked direction)
        // 4 moves = 16 points (completely open space)
        // We square the number of moves to create a stronger preference
        // for more open spaces
        totalScore += float64(moveCount * moveCount)
    }

    return totalScore
}
