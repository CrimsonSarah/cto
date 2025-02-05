package digidata

import "slices"

// Achieves fast deletion with amortized O(1) insertion by replacing
// an element with the last one on deletion. Order is therefore not
// preserved. On deletion, an O(n) search is performed and afterwards
// its deletion is done in O(1) time.
type UnorderedArray[T Comparable[T]] struct {
	Items []T
}

func MakeUnorderedArray[T Comparable[T]]() UnorderedArray[T] {
	return UnorderedArray[T]{}
}

func (a UnorderedArray[T]) Get() []T {
	return a.Items
}

func (a *UnorderedArray[T]) Add(element T) int {
	index := len(a.Items)
	a.Items = append(a.Items, element)

	return index
}

// Returns `true` if found an removed. `false` if not found.
func (a *UnorderedArray[T]) Remove(element T) bool {
	index := slices.IndexFunc(a.Items, func(candidate T) bool {
		return element.Equals(candidate)
	})

	if index == -1 {
		return false
	}

	if index < len(a.Items)-1 {
		last := a.Items[len(a.Items)-1]
		a.Items[index] = last
	}

	a.Items = a.Items[:len(a.Items)-1]
	return true
}

func (a *UnorderedArray[T]) Len() int {
	return len(a.Items)
}
