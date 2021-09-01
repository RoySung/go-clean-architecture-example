package utils

import (
	"fmt"
	"go-mock-example/domain"
)

func Bar(f domain.Foo) {
	result := f.Do(1)

	fmt.Println("result: ", result)
}
