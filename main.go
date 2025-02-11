package main

import (

	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "jacobstr",
		Color:      "#c35817",
		// Color:      "#4dcc8e",
		// Color:      "#3b210a",
		Head:       "default",
		Tail:       "train",
	}

	// portfolio := agent.NewPortfolio(
	// 	agent.NewHeuristic(4, "team-health", HeuristicHealth),
	// 	agent.NewHeuristic(1.7, "food", HeuristicFoodProximity),
	// 	agent.NewHeuristic(0.55, "Space to move around in", HeuristicMoveSpace),
	// 	agent.NewHeuristic(1.1, "Head to head", HeuristicAvoidCollisions),
	// 	agent.NewHeuristic(0.0009, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
	// )

	portfolio := agent.NewPortfolio(
		// agent.NewHeuristic(5, "team-health", HeuristicHealth),
		// agent.NewHeuristic(100, "food", HeuristicFoodProximity),
		// agent.NewHeuristic(100, "Space to move around in", HeuristicMoveSpace),
		// agent.NewHeuristic(170, "Self Instant Death", HeuristicInstantDeath),
		// agent.NewHeuristic(600, "Open Space", HeuristicOpenSpaceFloodFill),
		// agent.NewHeuristic(90, "Edge Discourage", HeuristicEdgeAvoidance),
		// agent.NewHeuristic(50, "Head to head", HeuristicHeadToHead),

		agent.NewHeuristic(1, "food", HeuristicFoodProximity),
		agent.NewHeuristic(1, "Space to move around in", HeuristicMoveSpace),
		agent.NewHeuristic(1, "Self Instant Death", HeuristicInstantDeath),
		agent.NewHeuristic(100, "Open Space", HeuristicOpenSpaceFloodFill),
		agent.NewHeuristic(1, "Edge Discourage", HeuristicEdgeAvoidance),
		agent.NewHeuristic(1, "Head to head", HeuristicHeadToHead),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
