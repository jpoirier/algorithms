/*

*/
package search


func BinarySearch(array []int, t int) bool {
	low := 0
	high := len(array) - 1
	for {
		if low <= high {
			ix := (low + high) / 2
			if t == array[ix] {
				return true
			} else if t < array[ix] {
				high = ix - 1
			} else {
				low = ix + 1
			}
			continue
		}
		break
	}
	return false
}
