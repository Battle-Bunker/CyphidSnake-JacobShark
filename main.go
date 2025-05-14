package main
// THis snake sucks, DO not attempt to work on it
// Please update this counter as a warnin g for the next dev:
// hours wasted = 50

import (

	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "jacobstr",
		Color:      "#FFFFFF",
		// Color:      "#a30303",
		// Color:      "#c35817",
		// Color:      "#4dcc8e",
		// Color:      "#3b210a",
		Head:       "default",
		Tail:       "replit-notmark",
	}

	// portfolio := agent.NewPortfolio(
	// 	agent.NewHeuristic(4, "team-health", HeuristicHealth),
	// 	agent.NewHeuristic(1.7, "food", HeuristicFoodProximity),
	// 	agent.NewHeuristic(0.55, "Space to move around in", HeuristicMoveSpace),
	// 	agent.NewHeuristic(1.1, "Head to head", HeuristicAvoidCollisions),
	// 	agent.NewHeuristic(0.0009, "Min A* for me max A* for enemy", HeuristicAStarFoodDistance),
	// )

	portfolio := agent.NewPortfolio(
		// agent.NewHeuristic(100, "food", HeuristicFoodProximity),
		// agent.NewHeuristic(100, "Space to move around in", HeuristicMoveSpace),
		// agent.NewHeuristic(170, "Self Instant Death", HeuristicInstantDeath),
		// agent.NewHeuristic(600, "Open Space", HeuristicOpenSpaceFloodFill),
		// agent.NewHeuristic(90, "Edge Discourage", HeuristicEdgeAvoidance),
		// agent.NewHeuristic(50, "Head to head", HeuristicHeadToHead),

		agent.NewHeuristic(5, "team-health", HeuristicHealth),
		agent.NewHeuristic(10, "food", HeuristicFoodProximity),
		// agent.NewHeuristic(5, "Space to move around in", HeuristicFloodFill),
		// agent.NewHeuristic(1, "Self Instant Death", HeuristicInstantDeath),
		agent.NewHeuristic(5, "Open Space", HeuristicFloodFill),
		agent.NewHeuristic(10, "Edge Discourage", HeuristicEdgeAvoidance),
		agent.NewHeuristic(5, "Head to head", HeuristicHeadToHead),
		agent.NewHeuristic(3, "Snake Density", HeuristicSnakeAntiCramming), //very high values, dont increase
		// agent.NewHeuristic(0.1, "Minimum Floodfill for closest enemy", MinFLoodfillForEnemy),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata,
	agent.WithTemperature(2), // default 5
	agent.WithPerformanceLogging(true))
	server := server.NewServer(snakeAgent)

	server.Start()
}
