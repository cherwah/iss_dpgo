package main

import (
	"decorator/time_deco"
	"fmt"
	"strconv"
)

func sum_it(arr []int) {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

/*
Decorates the sum_it() function with an
instrumenting capability.
*/
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Without Decoration")
	sum_it(arr)

	fmt.Println("----------------------")

	fmt.Println("With Decoration")
	time_deco.Time_it(sum_it, arr)
}
