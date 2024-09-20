package main

import (
	"fmt"
)

func SliceToChannel1(nums1 []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums1 {
			out <- n
		}
		close(out)
	}()
	return out
}

func SliceToChannel2(nums2 []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums2 {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	nums1 := []int{2, 3, 4, 7, 1}
	nums2 := []int{1, 5, 6, 8, 9}

	dataChannel1 := SliceToChannel1(nums1)
	finalChannel1 := sq(dataChannel1)

	dataChannel2 := SliceToChannel2(nums2)
	finalChannel2 := sq(dataChannel2)

	done := make(chan struct{})
	go func() {
		for n := range finalChannel1 {
			fmt.Println("channel 1:", n)
		}
		done <- struct{}{}
	}()

	go func() {
		for n := range finalChannel2 {
			fmt.Println("channel 2:", n)
		}
		done <- struct{}{}
	}()

	<-done
	<-done

}
