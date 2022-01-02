package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
)

var MAX float32 = 1000
var MIN float32 = -MAX

type GameState = [9]string

type Node struct {
	gameState GameState
	value     float32
	children  []Node
}

func minimaxRec(node *Node, player string, a float32, b float32, isMax bool, level float32) float32 {
	children := getNextMoves(node.gameState)

	if len(children) == 0 {
		winner := checkWinner(node.gameState)
		var value float32 = -1 / level
		if winner == player {
			value = 1 / level
		} else if winner == "draw" {
			value = 0
		}
		return value
	}

	childrenNodes := make([]Node, 0, 10)
	var value float32 = MAX
	if isMax {
		value *= MIN
	}

	for _, childState := range children {
		child := Node{gameState: childState}

		eval := minimaxRec(
			&child,
			player,
			a,
			b,
			!isMax,
			level+1,
		)

		child.value = eval

		if isMax && eval > value {
			value = eval
		} else if !isMax && eval < value {
			value = eval
		}

		childrenNodes = append(childrenNodes, child)

		if isMax {
			if value > a {
				a = value
			}
			if value >= b {
				break
			}
		} else {
			if value < b {
				b = value
			}
			if value <= a {
				break
			}
		}
	}

	node.children = childrenNodes
	return value
}

func minimax(gameState GameState) GameState {
	player := whosTurn(gameState)
	n := Node{gameState: gameState}
	value := minimaxRec(&n, player, MIN, MAX, true, 0)

	for _, child := range n.children {
		if child.value == value {
			return child.gameState
		}
	}

	return gameState
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func printBoard(gameState GameState) {
	winner := checkWinner(gameState)
	clearConsole()

	for i, cell := range gameState {
		if cell == "" && winner != "" {
			cell = " "
		} else if cell == "" {
			cell = color.InPurple(strconv.Itoa(i))
		}

		if (i+1)%3 == 0 {
			fmt.Print(cell + "\n")
		} else {
			fmt.Print(cell + " | ")
		}
	}

	fmt.Println()
}

func checkWinner(gameState GameState) string {
	winningCellCombinations := [...][3]int{
		// horizontal
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		// vertical
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		// diagnol
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combination := range winningCellCombinations {
		won := true
		player := gameState[combination[0]]

		for _, index := range combination {
			if gameState[index] != player {
				won = false
				break
			}
		}

		if won && player != "" {
			return player
		}
	}

	for _, cell := range gameState {
		if cell == "" {
			return ""
		}
	}

	return "draw"
}

func whosTurn(gameState GameState) string {
	moves := 0

	for _, cell := range gameState {
		if cell != "" {
			moves++
		}
	}

	if moves%2 == 0 {
		return "X"
	} else {
		return "O"
	}
}

func getNextMoves(gameState GameState) []GameState {
	a := make([]GameState, 0, 6)
	b := make([]GameState, 0, 6)

	if checkWinner(gameState) != "" {
		return a
	}

	player := whosTurn(gameState)

	for i, cell := range gameState {
		if cell == "" {
			clone := gameState
			clone[i] = player
			if rand.Intn(2) == 1 {
				a = append(a, clone)
			} else {
				b = append(b, clone)
			}
		}
	}

	return append(a, b...)
}

func move(gameState *GameState, index int) GameState {
	player := whosTurn(*gameState)
	gameState[index] = player
	return *gameState
}

func getMove(gameState *GameState) {
	fmt.Printf("Enter index for next move: ")
	var index int
	fmt.Scanln(&index)
	if index > 8 || gameState[index] != "" {
		fmt.Println("Illegal move")
		getMove(gameState)
	} else {
		move(gameState, index)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := GameState{
		"", "", "",
		"", "", "",
		"", "", "",
	}

	i := 0
	for checkWinner(game) == "" {
		if i%2 == 0 {
			game = minimax(game)
		} else {
			getMove(&game)
		}
		printBoard(game)
		i++
	}

	winner := checkWinner(game)
	if winner == "draw" {
		fmt.Println(winner)
	} else {
		fmt.Println(winner + " wins")
	}

}
