package pack

import (
	"fmt"
	"sort"
)

func Dedup() {
	// attrs := []string{"aa", "aa", "bb", "bb", "bb", "cc", "cc", "cc", "dd", "ee"}
	attrs := []string{"aa", "bb", "cc", "cc", "dd", "ee", "ee"}
	// attrs := []string{}
	sort.Strings(attrs)
	for i := len(attrs) - 1; i > 0; i-- {
		if attrs[i] == attrs[i-1] {
			attrs = append(attrs[:i], attrs[i+1:]...)
		}
	}
	fmt.Println(attrs)
}
