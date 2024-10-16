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
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}

	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(4, "team-health", HeuristicHealth),
		agent.NewHeuristic(1.2, "food", HeuristicFood),
		agent.NewHeuristic(0.9, "Space to move around in", HeuristicMoveSpace),
		agent.NewHeuristic(0.9, "Head to head", HeuristicAvoidCollisions),
		agent.NewHeuristic(0.0005, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
