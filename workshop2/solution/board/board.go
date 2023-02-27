package board

import (
	"fmt"
	"solution/asset"
)

type Board struct {
	length int
	width  int
	cell   [][]*asset.IPlayable
}

func New_Board(length int, width int) Board {
	board := Board{
		length: length,
		width:  width,
	}

	board.cell = make([][]*asset.IPlayable, length)

	for i := 0; i < length; i++ {
		board.cell[i] = make([]*asset.IPlayable, width)
	}

	return board
}

func (b *Board) Show() {
	for i := 0; i < b.length; i++ {
		str := ""

		for j := 0; j < b.width; j++ {
			if b.cell[i][j] != nil {
				var item asset.IPlayable = *b.cell[i][j]
				str += item.Get_Type() + " "
			} else {
				str += ". "
			}
		}

		fmt.Println(str)
	}
}

func (b *Board) Dimension() (int, int) {
	return b.length, b.width
}

func (b *Board) Set_Cell(playable asset.IPlayable, x int, y int) bool {
	if b.cell[x][y] == nil {
		b.cell[x][y] = &playable
		return true
	}

	return false
}
