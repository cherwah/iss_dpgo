package main

import (
	"flyweight/bitmap"
	"fmt"
)

type char struct {
	color     string
	is_bold   bool
	is_italic bool
	glyph     *bitmap.Bitmap_Font
}

type glyph_pool struct {
	glyphs [4]*bitmap.Bitmap_Font
}

func main() {
	// create a dynamic and empty array
	chars := []char{}

	// initialize bitmaps that represents our fonts
	glyph_pool := init_glyphs()

	// a test string of characters
	str := "0123001122231"

	// create character objects
	for _, s := range str {
		ch := new_char(string(s), &glyph_pool)
		chars = append(chars, ch)
	}

	// display our bitmap characters
	print_chars(chars)
}

func init_glyphs() glyph_pool {
	pool := glyph_pool{}

	pool.glyphs[0] = bitmap.Make_font0()
	pool.glyphs[1] = bitmap.Make_font1()
	pool.glyphs[2] = bitmap.Make_font2()
	pool.glyphs[3] = bitmap.Make_font3()

	return pool
}

func new_char(value string, g *glyph_pool) char {
	ch := char{
		color:     "#000000",
		is_bold:   false,
		is_italic: false,
	}

	if value == "0" {
		ch.glyph = g.glyphs[0]
	} else if value == "1" {
		ch.glyph = g.glyphs[1]
	} else if value == "2" {
		ch.glyph = g.glyphs[2]
	} else if value == "3" {
		ch.glyph = g.glyphs[3]
	}

	return ch
}

func print_chars(chars []char) {
	for i := 0; i < 5; i++ {
		for _, ch := range chars {
			bm := ch.glyph

			for j := 0; j < 5; j++ {
				fmt.Print(bm.Glyph[i][j])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
