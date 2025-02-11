package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
  _ "github.com/BattlesnakeOfficial/rules"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by Jacob)
func HeuristicHeadToHead(snapshot agent.GameSnapshot) float64 {
  var totalOverlap float64
  for _, snake := range snapshot.Snakes() {
    if snake.Length() < snapshot.You().Length() {
      continue
    }
    snakehead := snake.Head()
    snakeHeadCellNeighbours := snapshot.Board().Cells[snakehead.X][snakehead.Y].PassableNeighbours(snapshot.Board())



    mysnakehead := snapshot.You().Head()
    mysnakeHeadCellNeighbours := snapshot.Board().Cells[mysnakehead.X][mysnakehead.Y].PassableNeighbours(snapshot.Board())


    overlap := 0
    for _, neighbour := range snakeHeadCellNeighbours {
      for _, myneighbour := range mysnakeHeadCellNeighbours {
        if neighbour.Coordinates().X == myneighbour.Coordinates().X && neighbour.Coordinates().Y == myneighbour.Coordinates().Y {
          overlap++
        }
      }
    }
    totalOverlap += float64(overlap/4) * -10



  }
  return totalOverlap
}