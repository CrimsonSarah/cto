package digidata

// Used to speed up reads of the "all" kind.
type PrioritizedUnorderedArrayCache struct {
	IndexMap []int
	LenMap   []int
}

type PrioritizedUnorderedArray[T Comparable[T]] struct {
	Arrays []UnorderedArray[T]

	Length       int
	Cache        PrioritizedUnorderedArrayCache
	IsCacheValid bool
}

func MakePrioritizedUnorderedArray[T Comparable[T]]() PrioritizedUnorderedArray[T] {
	return PrioritizedUnorderedArray[T]{}
}

func (p *PrioritizedUnorderedArray[T]) Len() int {
	return p.Length
}

func (p *PrioritizedUnorderedArray[T]) Add(priority int, value T) {
	if priority >= len(p.Arrays) {
		diff := priority - len(p.Arrays) + 1
		p.Arrays = append(p.Arrays, make([]UnorderedArray[T], diff)...)
	}

	p.Arrays[priority].Add(value)
	p.IsCacheValid = false
	p.Length += 1
}

func (p *PrioritizedUnorderedArray[T]) RemoveP(priority int, value T) {
	removed := p.Arrays[priority].Remove(value)

	if removed {
		p.Length -= 1
	}
}

func (p *PrioritizedUnorderedArray[T]) Remove(value T) {
	for _, arr := range p.Arrays {
		removed := arr.Remove(value)

		if removed {
			p.Length -= 1
		}
	}
}

// You are NOT allowed to change the array while the cursor is being
// used :)
type PrioritizedUnorderedArrayCursor[T Comparable[T]] struct {
	Array *PrioritizedUnorderedArray[T]

	// Incremented on each Next().
	CurrentIndex int
	// Used in APIs that segregate each priority level.
	CurrentPriority int

	// The size should be equal to the amount of elements in all lists and
	// each value should be the index of the list in which it is contained.
	IndexMap []int
	// The cumulative length up to a given list, not counting itself.
	LenMap []int
}

func (p *PrioritizedUnorderedArray[T]) GetAll() PrioritizedUnorderedArrayCursor[T] {
	var indexMap, lenMap []int
	var accLen int

	if p.IsCacheValid {
		goto ret
	}

	// Cache miss.
	indexMap = make([]int, 0, len(p.Arrays))
	lenMap = make([]int, len(p.Arrays))
	accLen = 0

	for i, list := range p.Arrays {
		for range list.Len() {
			indexMap = append(indexMap, i)
		}

		lenMap[i] = accLen
		accLen += list.Len()
	}

	p.Cache.IndexMap = indexMap
	p.Cache.LenMap = lenMap
	p.IsCacheValid = true

ret:
	return PrioritizedUnorderedArrayCursor[T]{
		Array:           p,
		CurrentIndex:    0,
		CurrentPriority: 0,
		IndexMap:        p.Cache.IndexMap,
		LenMap:          p.Cache.LenMap,
	}
}

func (c *PrioritizedUnorderedArrayCursor[T]) Len() int {
	return len(c.IndexMap)
}

func (c *PrioritizedUnorderedArrayCursor[T]) IsEmpty() bool {
	return len(c.IndexMap) == 0
}

func (c *PrioritizedUnorderedArrayCursor[T]) Next() *T {
	if c.CurrentIndex >= len(c.IndexMap) {
		return nil
	}

	iOfArray := c.IndexMap[c.CurrentIndex]
	iInArray := c.CurrentIndex - c.LenMap[iOfArray]

	array := c.Array.Arrays[iOfArray]

	c.CurrentIndex += 1
	c.CurrentPriority = iOfArray

	// Will break if the array changes. But that is to be expected.
	return &array.Items[iInArray]
}

// Will not step over priority boundaries.
func (c *PrioritizedUnorderedArrayCursor[T]) NextInPriority() *T {
	if c.CurrentIndex >= len(c.IndexMap) {
		return nil
	}

	iOfArray := c.IndexMap[c.CurrentIndex]

	if iOfArray > c.CurrentPriority {
		return nil
	}

	iInArray := c.CurrentIndex - c.LenMap[iOfArray]

	array := c.Array.Arrays[iOfArray]
	c.CurrentIndex += 1

	// Will break if the array changes. But that is to be expected.
	return &array.Items[iInArray]
}

// This is expected to be used after NextInPriority returns nil.
// Returns false if there is no next priority.
func (c *PrioritizedUnorderedArrayCursor[T]) NextPriority() bool {
	nextPriority := c.CurrentPriority + 1

	// len(c.LenMap) should be the number of priorities.
	if nextPriority >= len(c.LenMap) {
		return false
	}

	c.CurrentPriority = nextPriority
	return true
}
