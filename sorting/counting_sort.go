/*

*/
package main

import (
	"fmt"
)

func CountingSort(array []int, k int) {
	l := len(array)
	b := make([]int, k)	// array of counts
	for i:= 0; i < l; i++ {
		b[array[i]] += 1
	}
	for i, j := 0, 0; i < k; i++ {
		for {
			if b[i] > 0 {
				array[j] = i
				j += 1
				b[i] -= 1
				continue
			}
			break
		}
	}
}

func main() {
	k := 25 // the array can contain values ranging from 0 to 24
	array := []int{19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 2}
	fmt.Println("Unsorted: ", array)
	CountingSort(array, k)
	fmt.Println("Sorted  : ", array)
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			fmt.Println("FAILED")
			return
		}
	}
	fmt.Println("PASSED")
}
