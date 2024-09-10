package main

import (
  "github.com/BattlesnakeOfficial/rules"
  "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicAvoidCollisions calculates a heuristic score based on potential collisions
func HeuristicAvoidCollisions(snapshot agent.GameSnapshot) float64 {
  you := snapshot.You()
  yourHead := you.Head()

  safeMoves := 0
  unsafeMoves := 0

  for _, move := range you.ForwardMoves() {
    var nextPos rules.Point
    switch move.Move {
    case "up":
      nextPos = rules.Point{X: yourHead.X, Y: yourHead.Y + 1}
    case "down":
      nextPos = rules.Point{X: yourHead.X, Y: yourHead.Y - 1}
    case "left":
      nextPos = rules.Point{X: yourHead.X - 1, Y: yourHead.Y}
    case "right":
      nextPos = rules.Point{X: yourHead.X + 1, Y: yourHead.Y}
    }

    if isUnsafeMove(nextPos, snapshot) {
      unsafeMoves++
    } else {
      safeMoves++
    }
  }

  // Return a score that rewards safe moves and penalizes unsafe ones
  return float64(safeMoves) - float64(unsafeMoves) * 10
}

func isUnsafeMove(pos rules.Point, snapshot agent.GameSnapshot) bool {
  for _, snake := range snapshot.Snakes() {
    if snake.ID() != snapshot.You().ID() {
      if isAdjacentTo(pos, snake.Head()) {
        return true
      }
    }
    for _, bodyPart := range snake.Body() {
      if pos == bodyPart {
        return true
      }
    }
  }
  return false
}

// Helper function to check if two points are adjacent
func isAdjacentTo(p1, p2 rules.Point) bool {
  dx := abs(int(p1.X - p2.X))
  dy := abs(int(p1.Y - p2.Y))
  return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

// Helper function for absolute value
func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}