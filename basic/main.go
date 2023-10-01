package main

import (
	"fmt"
)

// run command `go run basic` to run this program
func main() {
	intArr1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("original intArr1 capacity: %d, len: %d\n", cap(intArr1), len(intArr1))
	intArr1 = sliceDelete[int](intArr1, 3)
	fmt.Printf("intArr1: %v, cap: %d, len: %d\n", intArr1, cap(intArr1), len(intArr1))
	println("================== intArr1 end ==================")

	intArr2 := []int{}
	for i := 0; i < 32; i++ {
		intArr2 = append(intArr2, i)
	}
	fmt.Printf("original intArr2 capacity: %d, len: %d\n", cap(intArr2), len(intArr2))
	intArr2 = sliceDelete[int](intArr2, 31)
	fmt.Printf("intArr2: %v, cap: %d, len: %d\n", intArr2, cap(intArr2), len(intArr2))
	println("================== intArr2 end ==================")

	intArr3 := []int{}
	for i := 0; i < 33; i++ {
		intArr3 = append(intArr3, i)
	}
	fmt.Printf("original intArr3 capacity: %d, len: %d\n", cap(intArr3), len(intArr3))
	intArr3 = sliceDelete[int](intArr3, 0)
	fmt.Printf("intArr3: %v, cap: %d, len: %d\n", intArr3, cap(intArr3), len(intArr3))
}
