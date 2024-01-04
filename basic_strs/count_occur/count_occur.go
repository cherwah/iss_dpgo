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
  ch := ""

  fmt.Printf("Enter a long string: ")
  // create a reader that reads from standard input
  // (that is, from the command line)
  reader := bufio.NewReader(os.Stdin)

  // read until a newline character is found
  str, err := reader.ReadString('\n') 

  if err != nil {
    fmt.Println("Error reading input:", err)
    return
  }

  // trim the newline character as the reader
  // didn't filter the delimiter away
  str = strings.TrimRight(str, "\n")

  // only accepts a single character as input
  for ; ; {
    fmt.Printf("Enter character to count: ")
    fmt.Scanln(&ch)

    if len(ch) == 1 {
      // exit out of for-loop if a single character 
      // is provided
      break
    }

    // else, kept asking
    fmt.Printf("Please only enter a single character...\n")
  }

  fmt.Printf("'%s' appears %d times in your string.\n",
    ch, count_occurences(str, ch)) 
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


