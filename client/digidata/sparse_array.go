package digidata

// An array that is indexed by IDs that may not be continuous. IDs are
// not reutilized (unless it overflows or something).
type SparseArray[T any] struct {
	// We trade off access speed for the ability to not store data about
	// deleted elements.
	Indices map[int]int
	// We store them contiguously to increase the chance they will be
	// found in cache.
	Values []T
	// Used to find the ID of each index in the Values array. Useful for
	// when we need to reorder elements, for example when deleting.
	Ids []int

	IdCounter int
}

func MakeSparseArray[T any]() SparseArray[T] {
	return SparseArray[T]{
		Indices: make(map[int]int),
	}
}

func (a *SparseArray[T]) Len() int {
	return len(a.Values)
}

func (a *SparseArray[T]) Has(id int) bool {
	_, ok := a.Indices[id]
	return ok
}

func (a SparseArray[T]) Get(id int) (T, bool) {
	index, ok := a.Indices[id]

	if !ok {
		var empty T
		return empty, false
	}

	return a.Values[index], true
}

// BEWARE! The pointer will become invalid as soon as any operation is
// done on the array!
// Use with caution!
func (a SparseArray[T]) GetPtr(id int) (*T, bool) {
	index, ok := a.Indices[id]

	if !ok {
		var empty *T
		return empty, false
	}

	return &a.Values[index], true
}

func (a SparseArray[T]) GetAllUnordered() []T {
	return a.Values
}

func (a *SparseArray[T]) Update(id int, f func(element T) T) {
	index, ok := a.Indices[id]

	if !ok {
		return
	}

	a.Values[index] = f(a.Values[index])
}

func (a *SparseArray[T]) Remove(id int) {
	index, ok := a.Indices[id]

	if !ok {
		return
	}

	lastIndex := len(a.Values) - 1
	lastId := a.Ids[lastIndex]

	// Invalidate the ID of this element.
	delete(a.Indices, id)

	if index == lastIndex {
		goto chop
	}

	// Substitute said element for the last one.
	a.Values[index] = a.Values[lastIndex]
	// Substitute the ID of the last element appropriately.
	a.Ids[index] = a.Ids[lastIndex]

	// Update the last element's index in the Indices array.
	a.Indices[lastId] = index

chop:
	// Chop off the last element of each one.
	a.Values = a.Values[:lastIndex]
	a.Ids = a.Ids[:lastIndex]
}

func (a *SparseArray[T]) Add(element T) int {
	newId := a.IdCounter
	a.IdCounter += 1

	a.Indices[newId] = len(a.Values)
	a.Values = append(a.Values, element)
	a.Ids = append(a.Ids, newId)

	return newId
}

func (a *SparseArray[T]) AddWithId(f func(id int) T) int {
	newId := a.IdCounter
	a.IdCounter += 1

	a.Indices[newId] = len(a.Values)
	a.Values = append(a.Values, f(newId))
	a.Ids = append(a.Ids, newId)

	return newId
}

func (a *SparseArray[T]) Upsert(id int, element T) int {
	_, exists := a.Indices[id]

	if !exists {
		a.Indices[id] = len(a.Values)
		a.Values = append(a.Values, element)
		a.Ids = append(a.Ids, id)
	} else {
		a.Update(id, func(_ T) T { return element })
	}

	return id
}
