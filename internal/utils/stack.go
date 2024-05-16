package utils

type Container[T any] interface {
	Len() int
	Push(T)
	Pop() T
}

type Stack[T any] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	old := *s
	n := len(old)
	item := old[n-1]
	*s = (*s)[:len(*s)-1]
	return item
}
