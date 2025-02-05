package digidata

// Because `comparable` is only satisfied by builtin types.
type Comparable[T any] interface {
	Equals(T) bool
}
