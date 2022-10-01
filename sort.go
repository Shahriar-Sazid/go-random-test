package main

import (
	"fmt"
	"sort"
)

func sort_test() {
	a := []string{"b", "jfs", "d"}
	sort.Strings(a)
	fmt.Println(a)
}
