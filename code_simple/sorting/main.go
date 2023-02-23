package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// seed random generator
	rand.Seed(time.Now().UnixNano())

	// generate random array of integers
	rand_arr := array_random_int(10, 1, 50)
	fmt.Println("rand_arr =", rand_arr)

	// sort array
	insertion_sort(rand_arr)
	fmt.Println("sorted =", rand_arr)
}

/*
 * Returns an array of random integers.
 */
func array_random_int(size uint, min, max int) []int {
	var arr []int

	for i := uint(0); i < size; i++ {
		arr = append(arr, rand.Intn(max-min)+min)
	}

	return arr
}

/*
 * Sorting an array of integers using Insertion Sort algorithm.
 */
func insertion_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for pos := i; pos > 0; pos-- {
			if arr[pos] < arr[pos-1] {
				// swap values in 'pos' and 'pos-1'
				tmp := arr[pos-1]
				arr[pos-1] = arr[pos]
				arr[pos] = tmp
			} else {
				break
			}
		}
	}
}
