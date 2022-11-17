package pack

import "fmt"

func MapTest() {
	m := map[string]string{
		"bd":  "Bangladesh",
		"afg": "Afghanistan",
	}
	changeMap(m)
	fmt.Println("Map: ", m)
}

// map doesn't need to passed as pointer
func changeMap(m map[string]string) {
	m["pak"] = "Pakistan"
	delete(m, "bd")
}
