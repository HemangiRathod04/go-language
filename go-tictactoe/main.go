package main

import (
	"fmt"
	"strconv"
)

var board [3][3]string
var currentPlayer string

const (
	Red   = "\033[31m"
	Blue  = "\033[34m"
	Reset = "\033[0m"
)

func printCell(cell string) {
	switch cell {
	case "X":
		fmt.Print(Red + cell + Reset)
	case "O":
		fmt.Print(Blue + cell + Reset)
	default:
		fmt.Print(cell)
	}
}

func initGame() {
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = strconv.Itoa(count)
			count++
		}
	}
	currentPlayer = "X"
}

func printBoard() {
	fmt.Println("Current Board:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			printCell(board[i][j])
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("--|---|--")
		}
	}
	fmt.Println()
}

func checkWinner() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func checkTie() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != "X" && board[i][j] != "O" {
				return false
			}
		}
	}
	return true
}

func makeMove() bool {
	var position int
	fmt.Printf("Player %s, enter position (1-9): ", currentPlayer)
	fmt.Scanln(&position)

	if position >= 1 && position <= 9 {
		row, col := mapInputToPosition(position)
		if board[row][col] != "X" && board[row][col] != "O" {
			board[row][col] = currentPlayer
			return true
		} else {
			fmt.Println("Position already taken, try again.")
		}
	} else {
		fmt.Println("Invalid input, please enter a number between 1 and 9.")
	}
	return false
}

func mapInputToPosition(input int) (int, int) {
	row := (input - 1) / 3
	col := (input - 1) % 3
	return row, col
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to Tic-Tac-Toe!")
	printBoard()

	for {
		if makeMove() {
			printBoard()

			if checkWinner() {
				fmt.Printf("Player %s wins!\n", currentPlayer)
				break
			}

			if checkTie() {
				fmt.Println("It's a tie!")
				break
			}
			switchPlayer()
		}
	}
	fmt.Println("Game Over!")
}
