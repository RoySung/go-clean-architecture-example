package utils

import (
	"go-mock-example/domain"
	"go-mock-example/domain/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBar(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFoo := mocks.NewMockFoo(ctrl)
	mockFoo.EXPECT().Do(100).Return(100).Times(1)

	// Call.Do
	// mockFoo.EXPECT().Do(gomock.Any()).Do(func(i int) {
	// 	if i != 100 {
	// 		t.Errorf("parameter is not 100. get %d", i)
	// 	}
	// }).Return(100)

	// mockFoo2 := mocks.NewMockFoo2(ctrl)
	// mockFoo2.EXPECT().Do2(100).Return(100)

	type args struct {
		f domain.Foo
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"first",
			args{
				mockFoo,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoodBar(tt.args.f); got != tt.want {
				t.Errorf("Bar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBar(b *testing.B) {
	ctrl := gomock.NewController(b)
	mockFoo := mocks.NewMockFoo(ctrl)
	mockFoo.EXPECT().Do(100).Return(100).Times(b.N)

	for i := 0; i < b.N; i++ {
		GoodBar(mockFoo)
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"TEST ADD",
			args{1, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
