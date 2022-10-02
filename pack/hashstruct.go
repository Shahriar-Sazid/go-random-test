package pack

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

func HashStructTest() {
	type ComplexStruct struct {
		Name     string
		Age      *int
		Nums     []int
		Metadata map[string]interface{}
	}
	age1 := 64
	age2 := 64
	v1 := ComplexStruct{
		Name: "mitchellh",
		Age:  &age1,
		Nums: []int{3, 2, 1},
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}
	v2 := ComplexStruct{
		Name: "mitchellh",
		Age:  &age2,
		Nums: []int{3, 2, 1},
		Metadata: map[string]interface{}{
			"car":      true,
			"siblings": []string{"Bob", "John"},
			"location": "California",
		},
	}

	genAndPrint(v1)
	genAndPrint(v2)
}

func genAndPrint(v interface{}) {
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", hash)
}
