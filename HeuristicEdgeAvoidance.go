package main

import (
		_ "github.com/BattlesnakeOfficial/rules"
		 "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicEdgeAvoidance returns a score that penalizes snakes being near edges
// without explicitly rewarding central positions. The score is scaled to be
// roughly comparable to health-based heuristics for better integration with other
// heuristics in the portfolio.
func HeuristicEdgeAvoidance(snapshot agent.GameSnapshot) float64 {
		// score := 0.0
		// boardWidth := snapshot.Width()
		// boardHeight := snapshot.Height()

		// // We only care about our team's snakes
		// for _, snake := range snapshot.YourTeam() {
		// 		if !snake.Alive() {
		// 				continue
		// 		}

		// 		head := snake.Head()

		// 		// Calculate distance from each edge
		// 		distanceFromLeft := head.X
		// 		distanceFromRight := boardWidth - 1 - head.X
		// 		distanceFromTop := head.Y
		// 		distanceFromBottom := boardHeight - 1 - head.Y

		// 		// Apply penalty if within 2 spaces of any edge
		// 		// The closer to the edge, the higher the penalty
		// 		edgeThreshold := 2

		// 		if distanceFromLeft > edgeThreshold {
		// 				score -= float64(edgeThreshold - distanceFromLeft) * 10
		// 		}
		// 		if distanceFromRight > edgeThreshold {
		// 				score -= float64(edgeThreshold - distanceFromRight) * 10
		// 		}
		// 		if distanceFromTop > edgeThreshold {
		// 				score -= float64(edgeThreshold - distanceFromTop) * 10
		// 		}
		// 		if distanceFromBottom > edgeThreshold {
		// 				score -= float64(edgeThreshold - distanceFromBottom) * 10
		// 		}
		// }

		// return score
	snake := snapshot.You()
	// score := 0.0
	boardWidth := snapshot.Width()
	boardHeight := snapshot.Height()
	head := snake.Head()
	distanceFromLeft := head.X
	distanceFromRight := boardWidth - 1 - head.X
	distanceFromTop := head.Y
	distanceFromBottom := boardHeight - 1 - head.Y

	threshold := 1

	if distanceFromBottom < threshold {
		return float64(-20)
	}
	if distanceFromRight < threshold {
		return float64(-20)
	}
	if distanceFromTop < threshold {
		return float64(-20)
	}
	if distanceFromLeft < threshold {
		return float64(-20)
	}


	return 0
	
}
