package pack

import (
	"fmt"
)

func valueOrRef() {
	valueExample()
	refExample()
}

func valueExample() {
	fmt.Println("Value example")
	a := []int{1, 2, 3, 4, 5, 6}
	// a = nil //gives error
	b := a
	b[0] = -1
	// b = append(b, 7) //can't be appended
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
}

func refExample() {
	fmt.Println("Ref example")
	a := []int{1, 2, 3, 4, 5, 6}
	// a = nil // no error
	b := a
	b[0] = -1
	b = append(b, 7) //can be appended
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	fmt.Println("Capacity of a: ", cap(a))
	fmt.Println("Capacity of b: ", cap(b))
	b = append(b, 8)
	b = append(b, 9)
	b = append(b, 10)
	b = append(b, 11)
	b = append(b, 12)
	b = append(b, 13)
	fmt.Println("Length of b: ", len(b))
	fmt.Println("Capacity of b: ", cap(b))
	c := b[1:4]
	fmt.Println("Length of c: ", len(c))
	fmt.Println("Capacity of c: ", cap(c))
	b = append(b, 14)
	fmt.Println("Length of c: ", len(c))
	fmt.Println("Capacity of c: ", cap(c))
}
