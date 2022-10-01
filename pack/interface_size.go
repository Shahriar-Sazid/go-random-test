package pack

import (
	"fmt"
	"unsafe"
)

func interfaceSizeTest() {
	a := make([]interface{}, 10)
	a[0] = struct {
		A int
		B bool
	}{A: 1, B: false}

	a[1] = struct {
		C int
		D bool
	}{C: 1, D: false}

	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))
	fmt.Println("Allocated Memory: ", unsafe.Sizeof(a))
}
