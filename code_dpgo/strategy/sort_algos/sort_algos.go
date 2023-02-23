package sort_algos

/*
Interface that all our Sort algorithms will implement
*/
type Sorting interface {
	Sort(items []int)
}

/*
Sort types
*/
type BubbleSort struct {
}

type SelectionSort struct {
}

type InsertionSort struct {
}

/*
Different sort algorithms implementing the Sort interface
*/
func (algo BubbleSort) Sort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	// fmt.Print("After bubble-sort: ", arr, ".\n")
}

func (algo SelectionSort) Sort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		min_idx := i

		for j := i + 1; j < len; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}

		if min_idx != i {
			tmp := arr[i]
			arr[i] = arr[min_idx]
			arr[min_idx] = tmp
		}
	}
}

func (algo InsertionSort) Sort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
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
