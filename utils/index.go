package utils

import (
	"fmt"
	"go-mock-example/domain"
)

func Bar(f domain.Foo) int {
	result := f.Do(100)

	fmt.Println("result: ", result)
	return result
}
