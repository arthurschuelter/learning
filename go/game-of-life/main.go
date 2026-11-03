package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	cells [30][30]Cell
	size  int
}

type Cell struct {
	value      int
	neighbours int
}

func main() {
	size := 30
	board := startBoard(size)
	generations := 1000

	for gen := 0; gen < generations; gen++ {
		clearScreen()
		updateBoard(&board)
		printBoard(board, gen)
		pause()
	}
}

func updateBoard(board *Board) {
	verifyNeighbours(board)

	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			deadOrAlive(&board.cells[i][j])
		}
	}
}

func startBoard(size int) Board {
	var board Board
	board.size = size

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var random = rand.Float64()
			if random < 0.20 {
				board.cells[i][j].value = 1
			}
		}
	}

	return board
}

func verifyNeighbours(board *Board) {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			board.cells[i][j].neighbours = CountNeighbours(board, i, j)
		}
	}
}
func CountNeighbours(board *Board, i int, j int) int {
	count := 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			ni := i + x
			nj := j + y
			if ni >= 0 && ni < board.size && nj >= 0 && nj < board.size {
				count += board.cells[ni][nj].value
			}
		}
	}
	return count
}

func deadOrAlive(cell *Cell) {
	if cell.value == 1 {
		if cell.neighbours < 2 || cell.neighbours > 3 {
			cell.value = 0
		}
	} else {
		if cell.neighbours == 3 {
			cell.value = 1
		}
	}
	cell.neighbours = 0
}

func printBoard(board Board, generation int) {
	fmt.Printf("Generation: %d\n", generation)
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			var ch string
			if board.cells[i][j].value == 1 {
				ch = "■ "
			} else {
				ch = "  "
			}
			fmt.Print(ch)
		}
		fmt.Println()
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pause() {
	time.Sleep(100 * time.Millisecond)
}
