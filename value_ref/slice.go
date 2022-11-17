package vr

import "fmt"

func SliceTest() {
	a := []int{1, 2, 3}
	deleteSliceIndexWrong(a, 1)
	fmt.Println("array: ", a)
	deleteSliceIndexRight(&a, 1)
	fmt.Println("array: ", a)
}

func deleteSliceIndexRight(a *[]int, x int) {
	if x == len(*a)-1 {
		*a = (*a)[:len(*a)-1]
		return
	}
	*a = append((*a)[:x], (*a)[x+1:]...)
}

func deleteSliceIndexWrong(a []int, x int) {
	if x == len(a)-1 {
		a = (a)[:len(a)-1]
		return
	}
	a = append((a)[:x], (a)[x+1:]...)
}
