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

  // parameters to generate array of random numbers
  var n uint = 10
  min := 1
  max := 50

	// generate random array of integers
	rand_arr := array_random_int(n, min, max)
	fmt.Println("rand_arr =", rand_arr)

  target := rand.Intn(max-min+1)+min
  fmt.Printf("looking for %d...\n", target)

	// using Go's built-in sort()
	sort.Ints(rand_arr)

	// performs binary search
  found := binary_search(rand_arr, target)

  if found {
    fmt.Printf("%d is found.\n", target)
  } else {
    fmt.Printf("%d is not found.\n", target)
  }
}

/*
 * Returns an array of random integers.
 */
func array_random_int(size uint, min, max int) []int {
	var arr []int

	for i := uint(0); i < size; i++ {
		arr = append(arr, rand.Intn(max-min+1)+min)
	}

	return arr
}

/*
 * Locate the position of target in a sorted array
 * using Binary Search.
 *
 * Returns true if found; false otherwise.
 */
func binary_search(arr []int, target int) bool {
	start_pt := 0
	end_pt := len(arr) - 1
	mid_pt := (start_pt + end_pt) / 2

  found := false

	for { // while true
		if arr[mid_pt] == target {
      found = true
      break
		}

		if arr[mid_pt] > target {
			// search space should be further UP the array
			end_pt = mid_pt - 1
		} else {
			// search space should be further DOWN the array
			start_pt = mid_pt + 1
		}

		if start_pt <= end_pt {
		  // re-position our mid-point
		  mid_pt = (start_pt + end_pt) / 2
    } else {
      break
    }
	}

  return found
}
