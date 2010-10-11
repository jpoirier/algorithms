/*

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
