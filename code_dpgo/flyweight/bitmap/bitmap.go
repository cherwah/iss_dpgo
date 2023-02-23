package bitmap

type Bitmap_Font struct {
	Value string
	Glyph [5][5]string
}

func Make_font0() *Bitmap_Font {
	return &Bitmap_Font{
		Value: "0",
		Glyph: [5][5]string{
			{"#", "#", "#", "#", "#"},
			{"#", " ", " ", " ", "#"},
			{"#", " ", " ", " ", "#"},
			{"#", " ", " ", " ", "#"},
			{"#", "#", "#", "#", "#"},
		},
	}
}

func Make_font1() *Bitmap_Font {
	return &Bitmap_Font{
		Value: "1",
		Glyph: [5][5]string{
			{" ", " ", "#", " ", " "},
			{" ", " ", "#", " ", " "},
			{" ", " ", "#", " ", " "},
			{" ", " ", "#", " ", " "},
			{" ", " ", "#", " ", " "},
		},
	}
}

func Make_font2() *Bitmap_Font {
	return &Bitmap_Font{
		Value: "2",
		Glyph: [5][5]string{
			{"#", "#", "#", "#", "#"},
			{" ", " ", " ", " ", "#"},
			{"#", "#", "#", "#", "#"},
			{"#", " ", " ", " ", " "},
			{"#", "#", "#", "#", "#"},
		},
	}
}

func Make_font3() *Bitmap_Font {
	return &Bitmap_Font{
		Value: "3",
		Glyph: [5][5]string{
			{"#", "#", "#", "#", "#"},
			{" ", " ", " ", " ", "#"},
			{"#", "#", "#", "#", "#"},
			{" ", " ", " ", " ", "#"},
			{"#", "#", "#", "#", "#"},
		},
	}
}
