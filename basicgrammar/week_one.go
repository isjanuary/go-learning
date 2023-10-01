package main

import "fmt"

// 第一周作业, slice 的删除操作

// slice 的删除操作
func sliceDelete[T any](s []T, index int) []T {
	oldLen := len(s)
	oldCap := cap(s)
	if oldLen == 0 || oldCap == 0 {
		panic("cannot delete element from slice with length 0 or capacity 0")
	}

	if index < 0 || index > oldLen-1 {
		panic(fmt.Sprintf("index %d is out of range!\n", index))
	}

	newLen := oldLen - 1
	threshold := 32
	// if sliceCap is greater than threshold and newLen less than (3/4)*sliceCap, then shrink the capacity of underlying array and return a new slice.
	// otherwise, keep current capacity.
	newCap := oldCap * 3 / 4
	if oldCap > threshold && newLen <= newCap {
		newSlice := make([]T, newLen, newCap)
		for i := 0; i < index; i++ {
			newSlice[i] = s[i]
		}
		for i := index; i < newLen; i++ {
			newSlice[i] = s[i+1]
		}
		return newSlice
	}

	for i := index; i < newLen; i++ {
		s[i] = s[i+1]
	}
	s = s[0:newLen]

	return s
}
