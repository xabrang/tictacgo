package main

import (
	"fmt"
	"math/rand"
	"time"
)

var playerList = [2]string{"O", "X"}
var isDone = false
var player int = 0

func main() {
	board := [3][3]string{}
	turn := 0

	// randomize player
	rand.Seed(time.Now().UnixNano())
	player = rand.Intn(100) % 2

	fmt.Printf("TicTacGo : Playing as %s\n", playerList[player])

	for !isDone {
		move(&turn, &board)
		if checkBoard(&board) {
			printBoard(&board)
			fmt.Println("Tie!")
			isDone = true
		}
	}
}

func move(t *int, b *[3][3]string) {
	p := *t % 2
	input := [2]int{}

	printBoard(b)
	fmt.Printf("[%s] turn: ", playerList[p])

	if p != player {
		input = aiMove(*b)
	} else {
		fmt.Scanf("%d,%d", &input[0], &input[1])
	}

	if isValidMove(input, b) {
		b[input[0]][input[1]] = playerList[p]
		if isWinningMove(&playerList[p], &input, b) {
			fmt.Printf("%s win!\n", playerList[p])
			isDone = true
		}

		*t++
	}

}

func isValidMove(i [2]int, b *[3][3]string) bool {
	if i[0] > 2 || i[1] > 2 || i[0] < 0 || i[1] < 0 || (*b)[i[0]][i[1]] != "" {
		fmt.Println("Invalid Input!")
		return false
	}

	return true
}

// return True if no more move left
func checkBoard(b *[3][3]string) bool {
	for _, c := range *b {
		for _, s := range c {
			if s == "" {
				return false
			}
		}
	}
	return true
}

func isWinningMove(p *string, i *[2]int, b *[3][3]string) bool {
	w := true

	//check col
	for _, v := range (*b)[(*i)[0]] {
		if v != *p {
			w = false
		}
	}

	if w {
		return w
	}
	w = true

	//check row
	for _, xi := range *b {
		if xi[(*i)[1]] != *p {
			w = false
		}
	}
	if w {
		return w
	}

	w = false
	// check diagonal
	if (*b)[1][1] == *p {
		if (*b)[0][2] == *p && (*b)[0][2] == (*b)[2][0] {
			return true
		}
		if (*b)[0][0] == *p && (*b)[0][0] == (*b)[2][2] {
			return true
		}
	}

	return w
}

func printBoard(b *[3][3]string) {
	for _, c := range *b {
		for _, s := range c {
			if s == "" {
				s = " "
			}
			fmt.Printf("[%s]", s)
		}
		fmt.Println("")
	}
}

func aiMove(b [3][3]string) [2]int {
	bestMove := -1000
	move := [2]int{-1, -1}
	for x, c := range b {
		for y, s := range c {
			if s != "" {
				continue
			}
			b[x][y] = playerList[player]

			moveVal := minimax(b, 0, false)

			b[x][y] = ""

			if moveVal > bestMove {
				bestMove = moveVal
				move = [2]int{x, y}
			}
		}
	}
	fmt.Println(move[0], ",", move[1])
	return move
}

func evaluateMove(b *[3][3]string) int {
	// check row
	for x, _ := range *b {
		if (*b)[x][0] == (*b)[x][1] && (*b)[x][1] == (*b)[x][2] {
			if (*b)[x][0] == playerList[player] {
				return 10
			} else if (*b)[x][0] == playerList[(player+1)%2] {
				return -10
			}
		}
	}
	// check col
	for x, _ := range *b {
		if (*b)[0][x] == (*b)[1][x] && (*b)[1][x] == (*b)[2][x] {
			if (*b)[0][x] == playerList[player] {
				return 10
			} else if (*b)[0][x] == playerList[(player+1)%2] {
				return -10
			}
		}
	}
	// check diagonal
	if (*b)[0][0] == (*b)[1][1] && (*b)[1][1] == (*b)[2][2] {
		if (*b)[0][0] == playerList[player] {
			return 10
		} else if (*b)[0][0] == playerList[(player+1)%2] {
			return -10
		}
	}
	if (*b)[0][2] == (*b)[1][1] && (*b)[1][1] == (*b)[2][0] {
		if (*b)[1][1] == playerList[player] {
			return 10
		} else if (*b)[1][1] == playerList[(player+1)%2] {
			return -10
		}
	}
	return 0
}

func minimax(b [3][3]string, d int, isMax bool) int {
	score := evaluateMove(&b)
	best := -1000

	if score == 10 {
		return score
	}

	if score == -10 {
		return score
	}

	if checkBoard(&b) == false {
		return 0
	}

	if isMax {
		for x, c := range b {
			for y, s := range c {
				if s != "" {
					continue
				}

				b[x][y] = playerList[player]

				best = max(best, minimax(b, d+1, !isMax))

				b[x][y] = ""
			}
		}
		return best
	} else {
		for x, c := range b {
			for y, s := range c {
				if s != "" {
					continue
				}

				b[x][y] = playerList[player]

				best = min(best, minimax(b, d+1, !isMax))

				b[x][y] = ""
			}
		}

		return best
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
