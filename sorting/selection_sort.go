/*

*/
package sort


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
