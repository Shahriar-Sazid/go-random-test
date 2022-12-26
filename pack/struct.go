package pack

import "fmt"

type plusPlus struct {
	cnt int
}

func PlusPlusTest() {
	instance := plusPlus{}
	instance.cnt++
	fmt.Println(instance)
}
