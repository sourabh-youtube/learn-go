package generics

import "golang.org/x/exp/constraints"

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}
