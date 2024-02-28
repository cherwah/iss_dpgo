/*
  Author: Tan Cher Wah

  Go code to find the length of a string recursively.
*/

package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

/*
  Entry point.
*/
func main() {
  fmt.Printf("Enter string: ")
  
  // create a reader that reads from standard input
  // (that is, from the command line)
  reader := bufio.NewReader(os.Stdin)

  // read until a newline character is found
  str, err := reader.ReadString('\n')

  if err != nil {
    fmt.Println("Error reading input: ", err)
    return
  }

  // trim the newline character as the reader
  // didn't filter the delimiter away
  // "\r" to handle Windows newline character
  str = strings.TrimRight(str, "\r\n")

  fmt.Printf("Your string '%s' has %d characters.\n", 
    str, recursive_strlen(str))
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
