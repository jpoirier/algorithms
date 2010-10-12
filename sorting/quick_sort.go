/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package sort


func QuickSort(array []int) {
	if len(array) > 1 {
		// recursively work on the sub-arrays,
		// which are just slices of slices
		pivot := partition(array)
		QuickSort(array[0:pivot])
		QuickSort(array[pivot+1:])
	}
}
