package tictactoe

import (
	"fmt"
	"strings"
)

const (
	SYMBOL_PLAYER_1 = "X"
	SYMBOL_PLAYER_2 = "O"
	SYMBOL_EMPTY    = " "
)

var board = [3][3]string{}
var boardAsList = [9]string{}
var isPlayer1 = true

func StartGame() {
	resetGame()
	var position = 0
	for !isFull() {
		fmt.Printf("%s please select position (1-9) : ", getPlayerString())
		fmt.Scanf("%d", &position)
		fmt.Print("\n")
		valid, msg := checkPosition(position)
		if !valid {
			fmt.Printf("%s\n", msg)
		} else {
			fillPosition(position)
			printBoard()
			if isOver() {
				fmt.Printf("Congratulations %s you win", getPlayerString())
				break
			}
			isPlayer1 = !isPlayer1
		}
	}
}

func resetGame() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = SYMBOL_EMPTY
		}
	}

	for i := 0; i < len(boardAsList); i++ {
		boardAsList[i] = SYMBOL_EMPTY
	}
}

func printBoard() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (j + 1) == len(board[i]) {
				fmt.Printf("%s", board[i][j])
			} else {
				fmt.Printf("%s|", board[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func isFull() bool {
	for i := 0; i < len(boardAsList); i++ {
		if boardAsList[i] == SYMBOL_EMPTY {
			return false
		}
	}

	return true
}

func checkPosition(position int) (bool, string) {
	if position >= 1 && position <= 9 {
		if boardAsList[position-1] == SYMBOL_EMPTY {
			return true, ""
		}
		return false, "position is filled"
	}
	return false, "position is not valid"
}

func fillPosition(position int) {
	boardAsList[position-1] = getPlayerSymbol()
	var counter = 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			counter++
			if counter == position {
				board[i][j] = getPlayerSymbol()
				return
			}
		}
	}
}

func isOver() bool {
	symbol := getPlayerSymbol()
	return hasHorizontal(symbol) || hasVertical(symbol) || hasDiagonal(symbol)
}

func hasHorizontal(symbol string) bool {
	line := strings.Repeat(symbol, len(board))
	for i := 0; i < len(board); i++ {
		if strings.Join(board[i][:], "") == line {
			return true
		}
	}
	return false
}

func hasVertical(symbol string) bool {
	line := strings.Repeat(symbol, len(board))
	holder := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			holder += board[j][i]
		}
		if holder == line {
			return true
		} else {
			holder = ""
		}
	}
	return false
}

func hasDiagonal(symbol string) bool {
	line := strings.Repeat(symbol, len(board))
	holder1 := ""
	holder2 := ""
	for i := 0; i < len(board); i++ {
		holder1 += board[i][i]
		holder2 += board[(len(board)-1)-i][i]
	}

	if holder1 == line || holder2 == line {
		return true
	}

	return false
}

func getPlayerString() string {
	if isPlayer1 {
		return "Player 1"
	}
	return "Player 2"
}

func getPlayerSymbol() string {
	if isPlayer1 {
		return SYMBOL_PLAYER_1
	}
	return SYMBOL_PLAYER_2
}
