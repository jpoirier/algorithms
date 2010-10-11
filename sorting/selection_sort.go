/*

*/
package main

import (
	"fmt"
)

func selectMax(a []int) int {
	i := 0
	maxPos := 0
	right := len(a) - 1
	for {
		i += 1
		if i <= right {
			if a[i] > a[maxPos] {
				maxPos = i
			}
			continue
		}
		break
	}
	return maxPos
}

func SelectionSort(array []int) {
	for i := len(array)-1; i >= 1; i-- {
		maxPos := selectMax(array[0:i+1])
		if maxPos != i {
			array[i], array[maxPos] = array[maxPos], array[i]
		}
	}
}

func main() {
	array := []int{19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 2}
	fmt.Println("Unsorted: ", array)
	SelectionSort(array)
	fmt.Println("Sorted  : ", array)
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			fmt.Println("FAILED")
			return
		}
	}
	fmt.Println("PASSED")
}
