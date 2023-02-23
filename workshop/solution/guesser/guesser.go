package guesser

import "math/rand"

type IGuess interface {
	Guess() int
	Name() string
}

type Bounds struct {
	Min int
	Max int
}

type Next_guesser struct {
	Bounds
	N_tried int
}

type Rand_guesser struct {
	Bounds
}

func (g *Next_guesser) Guess() int {
	a_guess := g.Bounds.Min + g.N_tried
	g.N_tried += 1

	return a_guess
}

func (g *Next_guesser) Name() string {
	return "Next_guesser"
}

func (g *Rand_guesser) Guess() int {
	return rand.Intn(g.Bounds.Max-g.Bounds.Min) + g.Bounds.Min
}

func (g *Rand_guesser) Name() string {
	return "Rand_guesser"
}
