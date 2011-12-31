/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package sorting


type entry struct {
	element int
	next *entry
}

/* maintain count of entries in each bucket and pointer to its first entry */
type bucket struct {
	size int
	head *entry
}

var num int
var buckets []bucket

func hash(v int) int {
	return v/3
}

func extract(buckets []bucket, a []int) {
	idx := 0
	low := 0
	for i := 0; i < num; i++ {
		if buckets[i].size == 0 {
			continue
		}
		ptr := buckets[i].head
		if buckets[i].size == 1 {
			a[idx] = ptr.element
			idx += 1
			ptr = nil
			buckets[i].size = 0
			continue
		}
		low = idx
		a[idx] = ptr.element
		idx += 1
		ptr = ptr.next
		for {
			if ptr == nil {
				break
			}
			j := idx - 1
			for {
				if j >= low && a[j] > ptr.element {
					a[j + 1] = a[j]
					j -= 1
					continue
				}
				break
			}
			a[j + 1] = ptr.element
			ptr = ptr.next
			idx += 1
		}
		buckets[i].size = 0
	}
}

func BucketSort(array []int) {
	num = len(array)
	buckets = make([]bucket, num)
	for i := 0; i < num; i++ {
		k := hash(array[i])
		e := new(entry)
		e.element = array[i]
		if buckets[k].head == nil {
			buckets[k].head = e
		} else {
			e.next = buckets[k].head
			buckets[k].head = e
		}
		buckets[k].size += 1
	}
	extract(buckets, array)
}
