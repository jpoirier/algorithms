/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package sort


func selectKth(k int, array []int) int {
	pivot := partition(array)
	if k-1 == pivot {
		return pivot
	}
	if k-1 < pivot {
		return selectKth(k, array[0:pivot])
	}
	return selectKth(k-(pivot+1), array[pivot+1:])
}

func MedianSort(array []int) {
	l := len(array)
	if l > 1 {
		mid := l / 2
		selectKth(mid+1, array)
		// recursively work on the sub-arrays,
		// which are just slices of slices
		MedianSort(array[0:mid])
		MedianSort(array[mid:])
	}
}
