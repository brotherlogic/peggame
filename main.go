package main

import (
	"fmt"
)

type move struct {
	start  []int
	end    []int
	clears []int
}

type board struct {
	board [][]bool
	moves []*move
}

func initBoard() [][]bool {
	var board [][]bool
	board = make([][]bool, 5)
	for row := 0; row < 5; row++ {
		board[row] = make([]bool, row+1)
	}
	return board
}

func (b *board) printBoard() {
	for row := 0; row < 5; row++ {
		for i := 0; i < 4-row; i++ {
			fmt.Print(" ")
		}
		for _, val := range b.board[row] {
			if val {
				fmt.Print(". ")
			} else {
				fmt.Print("O ")
			}
		}

		fmt.Print("\n")
	}
}

func (b *board) getMoves() []*move {
	var moves []*move

	for row := 0; row < 5; row++ {
		for i, val := range b.board[row] {
			if val {
				// Vertical jump down
				if row == 4 && i == 2 {
					if !b.board[2][1] && !b.board[0][0] {
						moves = append(moves, &move{
							start:  []int{0, 0},
							end:    []int{4, 2},
							clears: []int{2, 1}})
					}
				}

				// Vertical jump up
				if row == 0 && i == 0 {
					if !b.board[2][1] && !b.board[4][2] {
						moves = append(moves, &move{
							end:    []int{0, 0},
							start:  []int{4, 2},
							clears: []int{2, 1}})
					}
				}

				// Horizontal jump left
				if i+2 <= row {
					if !b.board[row][i+1] && !b.board[row][i+2] {
						moves = append(moves, &move{
							start:  []int{row, i + 2},
							end:    []int{row, i},
							clears: []int{row, i + 1},
						})
					}
				}

				// Horizontal jump right
				if i > 1 {
					if !b.board[row][i-1] && !b.board[row][i-2] {
						moves = append(moves, &move{
							start:  []int{row, i - 2},
							end:    []int{row, i},
							clears: []int{row, i - 1},
						})
					}
				}

				// Diaganol jump down and left
				if row > 1 && i < row-1 {
					if !b.board[row-2][i] && !b.board[row-1][i] {
						moves = append(moves, &move{
							start:  []int{row - 2, i},
							end:    []int{row, i},
							clears: []int{row - 1, i},
						})
					}
				}

				// Diaganol jump down and right
				if row > 1 && i > 1 {
					if !b.board[row-2][i-2] && !b.board[row-1][i-1] {
						moves = append(moves, &move{
							start:  []int{row - 2, i - 2},
							end:    []int{row, i},
							clears: []int{row - 1, i - 1},
						})
					}
				}

				// Diaganol jump up and left {
				if row < 3 {
					if !b.board[row+2][i+2] && !b.board[row+1][i+1] {
						moves = append(moves, &move{
							start:  []int{row + 2, i + 2},
							end:    []int{row, i},
							clears: []int{row + 1, i + 1},
						})
					}
				}

				// Diaganol jump up and right
				if row < 3 {
					if !b.board[row+2][i] && !b.board[row+1][i] {
						moves = append(moves, &move{
							start:  []int{row + 2, i},
							end:    []int{row, i},
							clears: []int{row + 1, i},
						})
					}
				}
			}
		}
	}

	return moves
}

func (b *board) playMove(m *move) *board {
	nb := &board{}

	nb.board = initBoard()
	for i, row := range b.board {
		for j, e := range row {
			nb.board[i][j] = e
		}
	}

	nb.board[m.start[0]][m.start[1]] = true
	nb.board[m.end[0]][m.end[1]] = false
	nb.board[m.clears[0]][m.clears[1]] = true

	nb.moves = append(nb.moves, b.moves...)
	nb.moves = append(nb.moves, m)

	return nb
}

func (b *board) score() int {
	count := 0
	for row := 0; row < 5; row++ {
		for _, val := range b.board[row] {
			if !val {
				count++
			}
		}
	}
	return count
}

func (b *board) findBestSolution() *board {
	moves := b.getMoves()
	bestNum := 100
	var bestBoard *board
	for _, move := range moves {
		nb := b.playMove(move).findBestSolution()

		if nb.score() < bestNum {
			bestNum = nb.score()
			bestBoard = nb
		}
	}

	if bestBoard == nil {
		return b
	}

	return bestBoard
}

func main() {
	b := &board{
		board: initBoard(),
		moves: []*move{},
	}

	b.board[2][1] = true

	b.printBoard()
	fmt.Printf("Score: %v\n", b.score())
	fmt.Printf("Moves: %v\n\n", b.getMoves())

	s := b.findBestSolution()
	s.printBoard()
	fmt.Printf("Score: %v\n\n", s.score())

	fmt.Printf("Moves:\n")
	for _, m := range s.moves {
		fmt.Printf("Move: %+v\n", m)
	}
}
