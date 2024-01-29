package so

import "fmt"

type Test struct {
	Name string
}

func (t Test) String() string {
	return fmt.Sprintf("Name: %v", t.Name)
}

type StructWithName struct {
	Name string
}

type Constraint interface {
	StructWithName
	~struct {
		Name string
	}
	String() string
}

func Print[T Constraint](x T) {
	//s := T{}
	fmt.Printf("Hello %v", x)
}
