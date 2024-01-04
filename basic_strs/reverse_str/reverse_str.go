package main 

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {

  fmt.Printf("Enter string to reverse: ")


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

  fmt.Printf("You entered '%s'\n", str)
  fmt.Printf("The reverse is '%s'\n", reverse_str(str))
}

/*
 * Returns a string that has been reversed.
 */
func reverse_str(str string) string {
	var rstr = ""

  // reconstruct the string by reading it back
  // in reverse (start from the end)
	for i := len(str) - 1; i >= 0; i-- {
		// need to cast from a rune to a string
		rstr += string(str[i])
	}

	return rstr
}


