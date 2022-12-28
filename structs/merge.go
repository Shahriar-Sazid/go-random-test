package structs

import (
	"fmt"

	"github.com/imdario/mergo"
)

type foo struct {
	A string
	B int64
}

func MergeStructTest() {
	src := foo{
		A: "one",
		B: 2,
	}
	dest := foo{
		A: "two",
	}
	mergo.Merge(&dest, src)
	fmt.Println(dest)
	// Will print
	// {two 2}
}
