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
		aiMove(b)
	} else {
		fmt.Scanf("%d,%d", &input[0], &input[1])

		if isValidMove(input, b) {
			b[input[0]][input[1]] = playerList[p]
			if isWinningMove(&playerList[p], &input, b) {
				fmt.Printf("%s win!\n", playerList[p])
				isDone = true
			}

			*t++
		}
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

func aiMove(b *[3][3]string) {
	// TODO: Implements MiniMax
	s := ""
	fmt.Scanf("%s", &s)
}
