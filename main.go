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
		Color:      "#3b210a",
		Head:       "default",
		Tail:       "default",
	}

	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(4, "team-health", HeuristicHealth),
		agent.NewHeuristic(1.3, "food", HeuristicFoodProximity),
		agent.NewHeuristic(0.85, "Space to move around in", HeuristicMoveSpace),
		agent.NewHeuristic(1.1, "Head to head", HeuristicAvoidCollisions),
		agent.NewHeuristic(0.0009, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
