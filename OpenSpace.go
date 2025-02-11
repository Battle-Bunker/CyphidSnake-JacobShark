package main

import (
		"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// Point represents the coordinates on the board.
type Point = rules.Point

// Board represents the game board.
type Board = agent.Board

// Cell represents a single cell on the board.
type Cell = agent.Cell

// CellKind represents the type of a cell (e.g., Empty, Food, Snake).
type CellKind = agent.CellKind

// EmptyCell represents an empty cell on the board.
type EmptyCell = agent.EmptyCell

// FoodCell represents a food cell on the board.
type FoodCell = agent.FoodCell

// FloodFill performs a flood fill algorithm starting from (startX, startY) on the given board.
// It returns a slice of Points that are reachable from the starting position without collision.
func FloodFill(startX, startY int, board *Board) []Point {
		width := board.Width
		height := board.Height

		// Directions: up, down, left, right
		directions := []Point{
				{X: 0, Y: 1},
				{X: 1, Y: 0},
				{X: 0, Y: -1},
				{X: -1, Y: 0},
		}

		// Initialize visited matrix
		visited := make([][]bool, width)
		for i := range visited {
				visited[i] = make([]bool, height)
		}

		// Initialize queue for BFS
		type QueueItem struct {
				X, Y int
		}
		queue := []QueueItem{{X: startX, Y: startY}}
		visited[startX][startY] = true

		// Slice to store reachable cells
		reachable := []Point{{X: startX, Y: startY}}

		for len(queue) > 0 {
				// Dequeue the first element
				current := queue[0]
				queue = queue[1:]

				// Explore all four directions
				for _, dir := range directions {
						newX := current.X + dir.X
						newY := current.Y + dir.Y

						// Check bounds
						if newX < 0 || newX >= width || newY < 0 || newY >= height {
								continue
						}

						// Check if already visited
						if visited[newX][newY] {
								continue
						}

						// Get the cell at the new position
						cell := board.Cells[newX][newY]

						// Check if the cell is passable
						if cell.IsPassable() {
								visited[newX][newY] = true
								queue = append(queue, QueueItem{X: newX, Y: newY})
								reachable = append(reachable, Point{X: newX, Y: newY})
						}
				}
		}

		return reachable
}

// IsCellFree determines if a cell at (x, y) is free from walls and snake bodies.
func IsCellFree(x, y int, board *Board) bool {
		cell := board.Cells[x][y]
		return cell.IsPassable()
}

// AnalyzeTerritory analyzes the territory around the snake's head using Flood Fill.
func AnalyzeTerritory(board *Board, mySnake agent.SnakeSnapshot) []Point {
		head := mySnake.Head()
		return FloodFill(head.X, head.Y, board)
}