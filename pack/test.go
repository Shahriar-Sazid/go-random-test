package pack

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type name struct {
	name string
}

func mapToAnotherFunction(m map[string]name) {
	m["hello"] = name{name: "abc"}
	m["world"] = name{name: "abc"}
	m["new_word"] = name{name: "abc"}
}

func min[T constraints.Ordered](a T, b ...T) T {
	x := a
	for _, v := range b {
		if x > v {
			x = v
		}
	}
	return x
}

func Basic_test() {
	list := [10]name{}
	// list := make([]name, 0, 10)

	for i := 0; i < 10; i++ {
		list[i] = name{name: "abc"}
	}

	newList := list

	for ix := range newList {
		newList[ix].name = "def"
	}
	var slice = list[:]

	fmt.Printf("address of slice %p \n", slice)
	// fmt.Printf("address of slice %p \n", &slice)
	// fmt.Println(&slice[0])
	fmt.Printf("address of slice %p \n", list)
	fmt.Printf("address of slice %p \n", newList)

	// numList := []int{}
	numList := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		numList = append(numList, i)
	}

	newNumList := numList

	for ix := range newNumList {
		newNumList[ix] = newNumList[ix] * 2
	}

	fmt.Println(numList)
	fmt.Println(newNumList)

	// fmt.Println(list)

	// m := make(map[string]name)
	// m["hello"] = name{name: "abc"}
	// m["world"] = name{name: "def"}

	// // Initial State
	// for key, val := range m {
	// 	fmt.Println(key, "=>", val)
	// }

	// fmt.Println("-----------------------")

	// mapToAnotherFunction(m)
	// // After Passing to the function as a pointer
	// for key, val := range m {
	// 	fmt.Println(key, "=>", val)
	// }

	// var x = 5
	// var y = &x
	// fmt.Println(x, y)
	// var z = &y
	// fmt.Println(z)

	mp := map[uint]name{
		1: {name: "Sazid"},
		2: {name: "Fariha"},
		3: {name: "Ahmad"},
	}
	fariha := mp[1]
	fariha.name = "Fariha moni"

	fmt.Println(mp)

	abc := []uint{1, 2, 3, 4, 5}
	for _, v := range abc {
		switch v {
		case 1:
			fmt.Println(v)
		case 2:
			fmt.Println(v)
		case 3:
			fmt.Println(v)
		}
	}
	type FeatureListResponse struct {
		Features map[string]map[string]string `json:"features"`
		Profile  []string                     `json:"profile"`
	}

	fmt.Println(FeatureListResponse{})

	var list2 []int
	list2 = nil
	fmt.Println(list2)
	list2 = append(list2, 10)
	fmt.Println("abc")
	fmt.Println(list2)
	a := map[string]name{}
	a["abc"] = name{
		name: "sazid",
	}

	t := a["abc"]
	t.name = "ahmed"
	fmt.Println("changed", t)
	fmt.Println("changed", a)

	fmt.Println(min(1, 2, 3))
	fmt.Println(min(1))
}
