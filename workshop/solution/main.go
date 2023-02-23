package main

import (
	"fmt"
	"solution/guesser"
	"strconv"
)

func main() {
	/***************************************
	 * Set up boundaries and instances.
	 **************************************/
	bounds := guesser.Bounds{
		Min: 0,
		Max: 5,
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
	actual_val := 4
	var curr guesser.IGuess

	fmt.Println("Actual value: " + strconv.Itoa(actual_val))

	for {
		if turn%2 == 0 {
			curr = g1
		} else {
			curr = g2
		}

		the_guess := curr.Guess()
		fmt.Printf("%s guessed: %d\n", curr.Name(), the_guess)

		if the_guess == actual_val {
			fmt.Println(curr.Name() + " won!")
			break
		}

		turn++
	}
}
