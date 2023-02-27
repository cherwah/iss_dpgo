package main

import (
	"math/rand"
	"solution/asset"
	"solution/board"
)

func main() {
	// create game board
	board := board.New_Board(10, 10)

	length, width := board.Dimension()

	// add 5 humans to board
	for i := 0; i < 5; i++ {
		human := asset.New_Human()

		x := rand.Intn(length)
		y := rand.Intn(width)

		board.Set_Cell(human, x, y)
	}

	// add 5 monsters to board
	for i := 0; i < 5; i++ {
		monster := asset.New_Monster()

		x := rand.Intn(length)
		y := rand.Intn(width)

		board.Set_Cell(monster, x, y)
	}

	// show board with humans and monsters
	board.Show()
}
