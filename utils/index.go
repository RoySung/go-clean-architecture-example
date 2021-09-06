package utils

import (
	"fmt"
	"go-mock-example/domain"
)

type Foo struct{}

func (f Foo) Do(i int) int {
	return i
}

func BadBar() int {
	foo := Foo{}
	result := foo.Do(100)

	fmt.Println("result: ", result)
	return result
}

func GoodBar(f domain.Foo) int {
	result := f.Do(100)

	fmt.Println("result: ", result)
	return result
}

func Add(a int, b int) int {
	return a + b
}
