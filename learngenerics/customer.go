package learngenerics

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func GenericsOnStruct() {
	u := User[string]{
		ID:   1,
		Name: "John",
		Data: "description data",
	}

	fmt.Printf("USER with generic on structs: %#v\n", u)
}
