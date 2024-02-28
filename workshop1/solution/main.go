package main

import (
	"fmt"
	"solution/guesser"
)

func main() {
	/***************************************
	 * Set up boundaries and instances.
	 **************************************/
	bounds := guesser.Bounds{
		Min: 4,
		Max: 15,
	}

	g1 := &guesser.Next_guesser{
		Bounds:  bounds,
		N_tried: 0,
	}

	g2 := &guesser.Rand_guesser{
		Bounds: bounds,
	}

	/***************************************
	 * Play the guessing game.
	 **************************************/
	turn := 0
	correct_val := 11
	var curr guesser.IGuess

	fmt.Printf("Correct value: %d\n", correct_val)

	for {
		if turn%2 == 0 {
			curr = g1
		} else {
			curr = g2
		}

		the_guess := curr.Guess()
		fmt.Printf("%s guessed: %d\n", curr.Name(), the_guess)

		if the_guess == correct_val {
			fmt.Println(curr.Name() + " won!")
			break
		}

		turn++
	}
}
