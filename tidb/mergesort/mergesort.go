package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	tmp := doSort(src)
	for i := range tmp {
		src[i] = tmp[i]
	}
}

func doSort(src []int64) []int64 {

	num := len(src)

	if num == 1 {
		return src
	}

	mid := len(src) / 2

	var (
		left  = make([]int64, mid)
		right = make([]int64, num-mid)
	)
	for i := 0; i < num; i++ {
		if i < mid {
			left[i] = src[i]
		} else {
			right[i-mid] = src[i]
		}
	}

	return doMerge(doSort(left), doSort(right))
}

func doMerge(left, right []int64) []int64 {

	total := make([]int64, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {

		if left[0] < right[0] {
			total[i] = left[0]
			left = left[1:]
		} else {
			total[i] = right[0]
			right = right[1:]
		}

		i++
	}

	for j := range left {
		total[i] = left[j]
		i++
	}

	for j := range right {
		total[i] = right[j]
		i++
	}

	return total
}

func genRandSlice(size int) []int64 {

	r := make([]int64, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := range r {
		r[i] = rand.Int63n(200)
	}
	return r
}

func main() {

	src := genRandSlice(20)
	fmt.Println("Before", src)
	MergeSort(src)
	fmt.Println("After", src)

}
