/*
  Author: Tan Cher Wah

  Go code to sort an array of values.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {

	// seed random generator
	rand.Seed(time.Now().UnixNano())

  // create a reader that reads from standard input
  // (that is, from the command line)
  reader := bufio.NewReader(os.Stdin)

  // prompt user 
  fmt.Println("To randomly generate a list of values, " +
    "enter positive integers for NUM, LOW and HIGH " +
    " like so - 10 1 50")

  // read until a newline character is found
  str, err := reader.ReadString('\n') 

  if err != nil {
    fmt.Println("Error reading input: ", err)
    return
  }

  // trim the newline character as the reader
  // didn't filter the delimiter away
  str = strings.TrimRight(str, "\n")

  vec := []int{0, 0, 0}

  fdbk := parse_input(str, vec)
  if fdbk != "" {
    // print error message and exit
    fmt.Println(fdbk)
    os.Exit(1);
  }
  
	// generate random array of integers
	rand_arr := array_random_int(vec[0], vec[1], vec[2])
	fmt.Println("Generated Random Array =", rand_arr)

	// sort array
	insertion_sort(rand_arr)
	fmt.Println("Sorted =", rand_arr)
}

/*
 * Returns an array of random integers.
 */
func array_random_int(size int, min, max int) []int {
	var arr []int

	for i := 0; i < size; i++ {
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
        // item at 'pos' has found its place
        // in the sorted portion of the array
				break
			}
		}
	}
}

/*
 * Split input-string into individual parameters.
 */
func parse_input(str string, vec []int) string {
  parts := strings.Split(str, " ")

  // correct number of parameters?
  if len(parts) != 3 {
    return "Incorrect number of parameters provided."
  }

  // check that parameters are positive integers
  for i, part := range parts {
    // integer-check
    val, err := strconv.Atoi(strings.TrimSpace(part))
    vec[i] = val
    if err != nil {
      return fmt.Sprintf("Parameter %d is not an integer.", i+1)
    }

    // positive-check
    if vec[i] <= 0 {
      return fmt.Sprintf("Parameter %d is not positive.", i+1)
    }
  }

  // check that HIGH > LOW
  if vec[2] < vec[1] {
    return "LOW must be smaller than HIGH"
  }

  // no error
  return "" 
}





