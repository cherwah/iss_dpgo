/*
  Author: Tan Cher Wah
  
  Simple operations involving strings in Go.
*/

package main

import (
	"fmt"
)

/*
  Entry point.
*/
func main() {
	str := "hello golang!"
	fmt.Printf("'%s' reversed is '%s'\n", str, reverse_string(str))

	target := "g"
	fmt.Printf("There are %d occurrences of character '%s' in '%s'\n", 
    count_occurences(str, target), target, str)

	fmt.Println("There are", recursive_strlen(str),
		"characters in", str)
}

/*
 * Returns the number of occurrences of a character
 * within a string.
 */
func count_occurences(str string, target string) int {
	count := 0

	// 'range' provides an index and a value
	for _, ch := range str {
		if string(ch) == target {
			count++
		}
	}

	return count
}

/*
 * Returns a string that has been reversed.
 */
func reverse_string(str string) string {
	var rstr = ""

	for i := len(str) - 1; i >= 0; i-- {
		// need to cast from a rune to a string
		rstr += string(str[i])
	}

	return rstr
}

/*
 * Computes the length of a string using a recursive approach.
 *
 * Returns the length of the string.
 */
func recursive_strlen(str string) int {
	if len(str) == 0 {
		return 0
	}

	return 1 + recursive_strlen(str[1:])
}
