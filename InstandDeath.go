package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)


func HeuristicInstantDeath(snapshot agent.GameSnapshot) float64 {
    if snapshot.You().Health() == 0 {
        return float64(-500)
    }
    return 0
}
