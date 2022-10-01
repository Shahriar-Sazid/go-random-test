package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type S struct {
	K1 string
	K2 int
}

func hash_test() {
	sa := []S{{K1: "foo", K2: 1}, {K1: "bar", K2: 2}, {K1: "baz", K2: 3}}
	sb := []S{{K1: "baz", K2: 3}, {K1: "bar", K2: 2}, {K1: "foo", K2: 1}}
	sc := []S{}

	a := Hash(sa)
	b := Hash(sb)
	c := Hash(sc)

	fmt.Println(Compare(a, b))
	fmt.Println(Compare(a, c))
}

func Compare(a, b []byte) bool {
	a = append(a, b...)
	c := 0
	for _, x := range a {
		c ^= int(x)
	}
	return c == 0
}

func Hash(s []S) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}
