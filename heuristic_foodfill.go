package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicMoveSpace calculates the number of contiguous empty spaces
// available to the snake, which represents its freedom of movement.
func HeuristicMoveSpace(snapshot agent.GameSnapshot) float64 {
	// Get the board dimensions
	width := snapshot.Width()
	height := snapshot.Height()

	// Create a 2D grid to represent the board
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}

	// Mark occupied spaces
	markOccupiedSpaces(board, snapshot)

	// Find the head of our snake
	head := snapshot.You().Head()

	// Perform a flood fill from the head to count available spaces
	availableSpaces := floodFill(board, head.X, head.Y)

	return float64(availableSpaces)
}

// markOccupiedSpaces marks all occupied spaces on the board
func markOccupiedSpaces(board [][]bool, snapshot agent.GameSnapshot) {
	// Mark snake bodies
	for _, snake := range snapshot.Snakes() {
		for _, point := range snake.Body() {
			board[point.Y][point.X] = true
		}
	}

	// Mark hazards
	for _, hazard := range snapshot.Hazards() {
		board[hazard.Y][hazard.X] = true
	}
}

// floodFill performs a flood fill algorithm to count contiguous empty spaces
func floodFill(board [][]bool, x, y int) int {
	if x < 0 || y < 0 || y >= len(board) || x >= len(board[0]) || board[y][x] {
		return 0
	}

	board[y][x] = true
	count := 1

	// Recursively fill in all four directions
	count += floodFill(board, x+1, y)
	count += floodFill(board, x-1, y)
	count += floodFill(board, x, y+1)
	count += floodFill(board, x, y-1)

	return count
}
