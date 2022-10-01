package pack

import (
	"encoding/json"
	"fmt"
)

func MarshalTest1() {
	n1 := Name1{
		FirstName:  "Shahriar",
		SecondName: "Ahmad",
	}
	js, err := json.Marshal(n1)
	if err != nil {
		fmt.Println("error occurred in marshalling json")
		return
	}
	var t1 Name1
	err = json.Unmarshal(js, &t1)
	if err != nil {
		fmt.Println("error occurred in unmarshalling json")
		fmt.Printf("%+v", err)
		return
	}
}
