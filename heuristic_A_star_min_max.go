package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
  "github.com/BattlesnakeOfficial/rules"
  "container/heap"
  "math"
)

// Constants for A* algorithm
const (
  MaxIterations = 1000 // Prevent infinite loops
  MaxDistance   = 100.0 // Maximum distance to return if no path found
)

// Node represents a point on the grid with its f, g, and h scores
type Node struct {
  Point    rules.Point
  f, g, h  float64
  parent   *Node
  index    int // for use in the priority queue
}

// PriorityQueue implements heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
  n := len(*pq)
  node := x.(*Node)
  node.index = n
  *pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  node := old[n-1]
  old[n-1] = nil  // avoid memory leak
  node.index = -1 // for safety
  *pq = old[0 : n-1]
  return node
}

func aStarDistance(start, end rules.Point, snapshot agent.GameSnapshot) float64 {
  openList := make(PriorityQueue, 0)
  heap.Init(&openList)

  closedSet := make(map[rules.Point]bool)
  nodeMap := make(map[rules.Point]*Node)

  startNode := &Node{Point: start, g: 0, h: heuristic(start, end), f: 0}
  heap.Push(&openList, startNode)
  nodeMap[start] = startNode

  iterations := 0
  for openList.Len() > 0 && iterations < MaxIterations {
    iterations++
    current := heap.Pop(&openList).(*Node)

    if current.Point == end {
      return current.g
    }

    closedSet[current.Point] = true

    for _, neighbor := range getNeighbors(current.Point, snapshot) {
      if closedSet[neighbor] {
        continue
      }

      tentativeG := current.g + 1 // Assuming each step costs 1

      neighborNode, exists := nodeMap[neighbor]
      if !exists {
        neighborNode = &Node{Point: neighbor}
        nodeMap[neighbor] = neighborNode
      }

      if tentativeG < neighborNode.g || !exists {
        neighborNode.parent = current
        neighborNode.g = tentativeG
        neighborNode.h = heuristic(neighbor, end)
        neighborNode.f = neighborNode.g + neighborNode.h

        if !exists {
          heap.Push(&openList, neighborNode)
        } else {
          heap.Fix(&openList, neighborNode.index)
        }
      }
    }
  }

  // If no path is found or max iterations reached, return MaxDistance
  return MaxDistance
}

func heuristic(a, b rules.Point) float64 {
  dx := math.Abs(float64(a.X - b.X))
  dy := math.Abs(float64(a.Y - b.Y))
  return dx + dy
}

func getNeighbors(p rules.Point, snapshot agent.GameSnapshot) []rules.Point {
  neighbors := []rules.Point{
    {X: p.X + 1, Y: p.Y},
    {X: p.X - 1, Y: p.Y},
    {X: p.X, Y: p.Y + 1},
    {X: p.X, Y: p.Y - 1},
  }

  validNeighbors := make([]rules.Point, 0, 4)

  for _, neighbor := range neighbors {
    if isValidMove(neighbor, snapshot) {
      validNeighbors = append(validNeighbors, neighbor)
    }
  }

  return validNeighbors
}

func isValidMove(p rules.Point, snapshot agent.GameSnapshot) bool {
  // Check if the point is within the board boundaries
  if p.X < 0 || p.X >= snapshot.Width() || p.Y < 0 || p.Y >= snapshot.Height() {
    return false
  }

  // Check if the point collides with any snake
  for _, snake := range snapshot.Snakes() {
    for _, bodyPart := range snake.Body() {
      if p == bodyPart {
        return false
      }
    }
  }

  // Check if the point is in a hazard (you may want to allow this but with a higher cost)
  for _, hazard := range snapshot.Hazards() {
    if p == hazard {
      return false
    }
  }

  return true
}

// HeuristicAStarFoodDistance calculates a heuristic score based on A* distances to food
func HeuristicAStarFoodDistance(snapshot agent.GameSnapshot) float64 {
  ourScore := 0.0
  enemyScore := 0.0

  food := snapshot.Food()

  // Calculate score for our team
  for _, allySnake := range snapshot.YourTeam() {
    ourScore += nearestFoodDistance(allySnake.Head(), food, snapshot)
  }

  // Calculate score for enemy team
  for _, enemySnake := range snapshot.Opponents() {
    enemyScore += nearestFoodDistance(enemySnake.Head(), food, snapshot)
  }

  // We want to maximize our distance and minimize enemy distance
  score := ourScore - enemyScore

  // Clamp the score between -100 and 100
  return math.Max(-100, math.Min(100, score))
}

// nearestFoodDistance calculates the A* distance to the nearest food
func nearestFoodDistance(start rules.Point, food []rules.Point, snapshot agent.GameSnapshot) float64 {
  if len(food) == 0 {
    return 0
  }

  minDistance := MaxDistance
  for _, f := range food {
    distance := aStarDistance(start, f, snapshot)
    if distance < minDistance {
      minDistance = distance
    }
  }

  return minDistance
}