package domain

//go:generate mockgen -source=foo.go -destination=mocks/mock_foo.go -package=mocks go-mock-example/domain/mocks Foo
type Foo interface {
	Do(int) int
}

type Foo2 interface {
	Do2(int) int
}
