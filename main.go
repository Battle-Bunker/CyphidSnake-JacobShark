package main

import (

	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "",
		Color:      "#C35817",
		// Color:      "#3b210a",
		Head:       "default",
		Tail:       "default",
	}

	// portfolio := agent.NewPortfolio(
	// 	agent.NewHeuristic(4, "team-health", HeuristicHealth),
	// 	agent.NewHeuristic(1.7, "food", HeuristicFoodProximity),
	// 	agent.NewHeuristic(0.55, "Space to move around in", HeuristicMoveSpace),
	// 	agent.NewHeuristic(1.1, "Head to head", HeuristicAvoidCollisions),
	// 	agent.NewHeuristic(0.0009, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
	// )

	portfolio := agent.NewPortfolio(
		// agent.NewHeuristic(2, "team-health", HeuristicHealth),
		agent.NewHeuristic(50, "food", HeuristicFoodProximity),
		agent.NewHeuristic(100, "Space to move around in", HeuristicMoveSpace),
		agent.NewHeuristic(70, "Head to head", HeuristicAvoidCollisions),
		// agent.NewHeuristic(0, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
		// agent.NewHeuristic(1, "Corner", HeuristicCorner),
		// agent.NewHeuristic(5, "Self Instant Food", HeuristicSelfDeath),
		agent.NewHeuristic(170, "Self Instant Death", HeuristicInstantDeath),
		// agent.NewHeuristic(-70, "Open Space", HeuristicOpenSpace),
		agent.NewHeuristic(90, "Edge Discourage", HeuristicEdgeAvoidance),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
