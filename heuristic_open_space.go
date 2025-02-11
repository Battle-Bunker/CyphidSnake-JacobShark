
package main

import (
    "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicOpenSpaceFloodFill calculates a score based on the amount of open space
// available around each snake using flood fill
func HeuristicOpenSpaceFloodFill(snapshot agent.GameSnapshot) float64 {
    board := snapshot.Board()
    yourSnake := snapshot.You()
    
    if !yourSnake.Alive() {
        return 0
    }
    
    // Get reachable spaces for your snake
    reachableSpaces := AnalyzeTerritory(board, yourSnake)
    
    // Return a score based on the number of reachable spaces
    return float64(len(reachableSpaces))
}
