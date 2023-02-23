package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// seed random generator
	rand.Seed(time.Now().UnixNano())

	// generate random array of integers
	rand_arr := array_random_int(10, 1, 50)
	fmt.Println("rand_arr =", rand_arr)

	// using Go's built-in sort()
	sort.Ints(rand_arr)

	// performs binary search
	var pos = binary_search(rand_arr, rand_arr[5])

	if pos != -1 {
		fmt.Println(rand_arr[5], "is found at pos", pos)
	} else {
		fmt.Println("Not found.")
	}
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
 * Locate the position of target in a sorted array
 * using Binary Search.
 *
 * Returns zero-index position of target in the array
 * if found; -1 otherwise.
 */
func binary_search(arr []int, target int) int {
	start_pt := 0
	end_pt := len(arr) - 1
	mid_pt := (start_pt + end_pt) / 2

	for { // while true
		if arr[mid_pt] == target {
			// found
			return mid_pt
		}

		if arr[mid_pt] > target {
			// search space should be further UP the array
			end_pt = mid_pt - 1
		} else {
			// search space should be further DOWN the array
			start_pt = mid_pt + 1
		}

		if start_pt > end_pt {
			break
		}

		// re-position our mid-point
		mid_pt = (start_pt + end_pt) / 2
	}

	// not found
	return -1
}
