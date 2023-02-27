package main

import (
	"math/rand"
	"solution/asset"
	"solution/board"
	"time"
)

func main() {
	// seed random generator
	rand.Seed(time.Now().UnixNano())

	// create game board
	board := board.New_Board(10, 10)
	length, width := board.Dimension()

	// add 5 humans to board
	for i := 0; i < 5; i++ {
		human := asset.New_Human()

		for {
			x := rand.Intn(length)
			y := rand.Intn(width)

			if board.Set_Cell(human, x, y) {
				break
			}
		}
	}

	// add 5 monsters to board
	for i := 0; i < 5; i++ {
		monster := asset.New_Monster()

		for {
			x := rand.Intn(length)
			y := rand.Intn(width)

			if board.Set_Cell(monster, x, y) {
				break
			}
		}
	}

	// show board with humans and monsters
	board.Show()
}
