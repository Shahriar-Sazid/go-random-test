package pack

import (
	"encoding/json"
	"fmt"
)

func MapJSON() {
	type mapJSON map[string]interface{}

	a := mapJSON{
		"Count":      int(1),
		"FloatCount": 2.43,
	}

	jsn, _ := json.Marshal(a)

	var b mapJSON

	_ = json.Unmarshal(jsn, &b)

	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", b["Count"], b["Count"]) // This is a float now not an integer
	fmt.Printf("%v %T\n", b["FloatCount"], b["FloatCount"])
	ab := append([]byte("hello "), "world"...)
	fmt.Println("ab: ", string(ab))
}
