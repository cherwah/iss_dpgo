package time_deco

import (
	"fmt"
	"strconv"
	"time"
)

// a simple decorator to track time-elasped
// for a function
func Time_it(f func([]int), arr []int) {
	start := time.Now()
	f(arr)
	fmt.Println("Took: " +
		strconv.Itoa(int(time.Since(start).Microseconds())) +
		" micro-seconds")
}
