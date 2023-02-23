package main

import (
	"fmt"
	"strategy/sort_algos"
)

func main() {
	// init different sorting strategies
	algo1 := sort_algos.BubbleSort{}
	algo2 := sort_algos.SelectionSort{}
	algo3 := sort_algos.InsertionSort{}

	/*
	 giving the same function different
	 sorting strategies
	*/
	arr1 := []int{3, 5, 1, 9, 4}
	do_work(arr1, algo1)

	arr2 := []int{7, 9, 5, 9, 2}
	do_work(arr2, algo2)

	arr3 := []int{0, 1, 9, 25, 15}
	do_work(arr3, algo3)
}

func do_work(arr []int, algo sort_algos.Sorting) {
	algo.Sort(arr)
	fmt.Print("After sort: ", arr, "\n")
}
