/*

*/
package main

import (
	"fmt"
)

func insert(a []int, pos int, val int) {
	i := pos - 1
	for {
		if i >= 0 && a[i] > val {
			a[i + 1] = a[i]
			i -= 1
			continue
		}
		a[i + 1] = val
		break
	}
}

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		insert(array, i, array[i])
	}
}

func main() {
	array := []int{88, 54, 37, 28, 9, 2, 1}
	fmt.Println("Unsorted: ", array)
	InsertionSort(array)
	fmt.Println("Sorted  : ", array)
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			fmt.Println("FAILED")
			return
		}
	}
	fmt.Println("PASSED")
}
