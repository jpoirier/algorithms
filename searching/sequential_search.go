/*

*/
package search


func SequentialSearch(array []int, t int) bool {
	for i := len(array)-1; i >= 1; i-- {
		if array[i] == t {
			return true
		}
	}
	return false
}
