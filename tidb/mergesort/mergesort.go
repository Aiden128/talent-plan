package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count = 0

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	res := make(chan struct{}, 4)
	tmp := goDoSemSort(src, res)
	for i := range tmp {
		src[i] = tmp[i]
	}
}

func goDoSemSort(src []int64, sem chan struct{}) []int64 {

	num := len(src)

	if num < 2 {
		return src
	}
	mid := num / 2

	pool := sync.WaitGroup{}
	pool.Add(2)

	var (
		left  = []int64{}
		right = []int64{}
	)

	select {
	case sem <- struct{}{}:
		go func() {
			left = goDoSemSort(src[:mid], sem)
			<-sem
			pool.Done()
		}()
	default:
		left = goDoSemSort(src[:mid], sem)
		pool.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			right = goDoSemSort(src[mid:], sem)
			<-sem
			pool.Done()
		}()
	default:
		right = goDoSemSort(src[mid:], sem)
		pool.Done()
	}

	pool.Wait()
	return doMerge(left, right)

}

func goDoSort(src []int64, res chan []int64) {

	num := len(src)

	if len(src) < 2 {
		res <- src
		return
	}

	mid := num / 2

	var (
		leftChan  = make(chan []int64)
		rightChan = make(chan []int64)
	)

	go goDoSort(src[mid:], leftChan)
	go goDoSort(src[:mid], rightChan)

	l := <-leftChan
	r := <-rightChan

	close(leftChan)
	close(rightChan)

	res <- doMerge(l, r)
	return
}

func doSort(src []int64) []int64 {

	num := len(src)

	if num == 1 {
		return src
	}
	mid := len(src) / 2
	var (
		left  = src[mid:]
		right = src[:mid]
	)

	return doMerge(doSort(left), doSort(right))
}

func doMerge(left, right []int64) []int64 {

	fmt.Println(count)
	count++

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
