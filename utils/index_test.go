package utils

import (
	"go-mock-example/domain"
	"testing"
)

func TestBar(t *testing.T) {
	type args struct {
		f domain.Foo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bar(tt.args.f)
		})
	}
}
