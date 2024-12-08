package primecheck

import "fmt"

type Set[T comparable] struct {
	Keys map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		Keys: make(map[T]struct{}),
	}
}

func (set *Set[T]) Add(value T) {
	set.Keys[value] = struct{}{}
}

func (set *Set[T]) Has(value T) bool {
	_, ok := set.Keys[value]
	return ok
}

func (set *Set[T]) Remove(value T) {
	delete(set.Keys, value)
}

func (set *Set[T]) PopAny() (ans T) {
	for key := range set.Keys {
		set.Remove(key)
		return key
	}

	return ans
}

func (set *Set[T]) IsEmpty() bool {
	return len(set.Keys) == 0
}

func (set *Set[T]) String() string {
	return fmt.Sprintf("%v", set.Keys)
}
