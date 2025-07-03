package gheap

type (
	// This allows calculation of either integer or float scores
	Number interface {
		~int | ~int32 | ~int64 | ~float32 | ~float64
	}

	HeapElement[S Number, T any] struct {
		Key  S
		Data T
	}

	MaxHeap[S Number, T any]    []HeapElement[S, T]
	MinHeap[S Number, T any]    []HeapElement[S, T]
	MinMaxHeap[S Number, T any] []HeapElement[S, T]
)

func GreaterThan[S Number](a, b S) bool {
	return a > b
}

func LessThan[S Number](a, b S) bool {
	return a < b
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

// *****************************************
// Below are functions that implement MinHeap
// *****************************************

func (b MinHeap[S, T]) shiftUp(i int) {
	p := parent(i)
	for i > 0 && b[p].Key > b[i].Key {
		b[p], b[i] = b[i], b[p]
		i = p
		p = parent(i)
	}
}

func (b MinHeap[S, T]) shiftDown(i int) {
	for {
		maxIndex := i
		l := leftChild(i)
		if l < len(b) && b[l].Key < b[maxIndex].Key {
			maxIndex = l
		}
		r := rightChild(i)
		if r < len(b) && b[r].Key < b[maxIndex].Key {
			maxIndex = r
		}
		if i == maxIndex { // If i is maxIndex, we are done!
			break
		}
		b[i], b[maxIndex] = b[maxIndex], b[i]
		i = maxIndex
	}
}

// Insert adds element to the heap and maintains the
// min heap order
func (b *MinHeap[S, T]) Insert(element HeapElement[S, T]) {
	*b = append(*b, element)
	b.shiftUp(len(*b) - 1) // Maintain the heap property
}

// RemoveMin returns the element with the highest key
// and removes it from the heap.
func (b *MinHeap[S, T]) RemoveMin() (result HeapElement[S, T]) {
	result = (*b)[0]
	(*b)[0] = (*b)[len(*b)-1] // Replace the value at the root  with the last leaf
	*b = (*b)[:len(*b)-1]     // Shorten slice by one
	b.shiftDown(0)            // Maintain the heap property
	return
}

// PeekMain returns the element with the lowest key
// without removing it.
func (b MinHeap[S, T]) PeekMin() HeapElement[S, T] {
	return b[0]
}

// Copy returns a shallow copy of the heap
func (b MinHeap[S, T]) Copy() (c MinHeap[S, T]) {
	c = append([]HeapElement[S, T]{}, b...)
	return
}

// *****************************************
// Below are functions that implement MaxHeap
// *****************************************

func (b MaxHeap[S, T]) shiftUp(i int) {
	p := parent(i)
	for i > 0 && b[p].Key < b[i].Key {
		b[p], b[i] = b[i], b[p]
		i = p
		p = parent(i)
	}
}

func (b MaxHeap[S, T]) shiftDown(i int) {
	for {
		maxIndex := i
		l := leftChild(i)
		if l < len(b) && b[l].Key > b[maxIndex].Key {
			maxIndex = l
		}
		r := rightChild(i)
		if r < len(b) && b[r].Key > b[maxIndex].Key {
			maxIndex = r
		}
		if i == maxIndex { // If i is maxIndex, we are done!
			break
		}
		b[i], b[maxIndex] = b[maxIndex], b[i]
		i = maxIndex
	}
}

// Insert adds element to the heap and maintains the
// max heap order
func (b *MaxHeap[S, T]) Insert(element HeapElement[S, T]) {
	*b = append(*b, element)
	b.shiftUp(len(*b) - 1) // Maintain the heap property
}

// PeekMax returns the element with the highest key
// without removing it.
func (b MaxHeap[S, T]) PeekMax() HeapElement[S, T] {
	return b[0]
}

// RemoveMax returns the element with the highest key
// and removes it from the heap.
func (b *MaxHeap[S, T]) RemoveMax() (result HeapElement[S, T]) {
	result = (*b)[0]
	(*b)[0] = (*b)[len(*b)-1] // Replace the value at the root  with the last leaf
	*b = (*b)[:len(*b)-1]     // Shorten slice by one
	b.shiftDown(0)            // Maintain the heap property
	return
}

// Copy returns a shallow copy of the heap
func (b MaxHeap[S, T]) Copy() (c MaxHeap[S, T]) {
	c = append([]HeapElement[S, T]{}, b...)
	return
}

// *****************************************
// Below are functions that implement MinMaxHeap
// *****************************************

// IsMinLevel returns true if and only if i is the index
// of an element in a minimum level of the MinMaxHeap
func IsMinLevel(i int) (isMin bool) {
	i++
	for i > 0 {
		i >>= 1
		isMin = !isMin
	}
	return
}

// pushGranny takes an index i and if it has a grand parent to which it
// compares affirmatively to the key of i, the elements are swapped
// and the process is repeated until there is no grand parent.
func (h MinMaxHeap[S, T]) pushGranny(i int, compare func(a, b S) bool) {
	for i > 2 { // Every index greater than 2 has a grand parent
		gp := parent(parent(i))
		if compare(h[i].Key, h[gp].Key) {
			h[i], h[gp] = h[gp], h[i]
		}
		i = gp
	}
}

func (h MinMaxHeap[S, T]) pushUp(i int) {
	if i > 0 {
		p := parent(i)
		if IsMinLevel(i) {
			if h[i].Key > h[p].Key {
				h[i], h[p] = h[p], h[i]
				h.pushGranny(p, GreaterThan[S])
			} else {
				h.pushGranny(i, LessThan[S])
			}
		} else {
			if h[i].Key < h[p].Key {
				h[i], h[p] = h[p], h[i]
				h.pushGranny(p, LessThan[S])
			} else {
				h.pushGranny(i, GreaterThan[S])
			}
		}
	}
}

// pushDown maintains heap order following a removal operation
// the compare function is either greaterThan or lessThan
func (h MinMaxHeap[S, T]) pushDown(i int, compare func(a, b S) bool) {
	for {
		l := leftChild(i)
		if l >= len(h) { // No Children for this node, we are done.
			break
		}
		minIndex := l
		r := rightChild(i)
		// The children and grand children indices of i that need
		// to be tested are arranged in ascending order.
		for _, v := range [5]int{r, leftChild(l), rightChild(l), leftChild(r), rightChild(r)} {
			if v >= len(h) {
				break // No kids beyond this point
				//continue
			}
			if compare(h[v].Key, h[minIndex].Key) {
				minIndex = v
			}
		}
		// Now that minIndex is identified, compare to i.
		if !compare(h[minIndex].Key, h[i].Key) {
			break
		}
		h[minIndex], h[i] = h[i], h[minIndex]
		if minIndex <= r { // minIndex is a child of i
			break
		}
		p := parent(minIndex)
		if compare(h[p].Key, h[minIndex].Key) {
			h[minIndex], h[p] = h[p], h[minIndex]
		}
		i = minIndex
	}
}

// Insert adds element to the heap and maintains the
// minmax heap order.
func (b *MinMaxHeap[S, T]) Insert(element HeapElement[S, T]) {
	*b = append(*b, element)
	b.pushUp(len(*b) - 1)
}

// RemoveMin returns the element with the lowest key
// and removes it from the heap.
func (b *MinMaxHeap[S, T]) RemoveMin() (result HeapElement[S, T]) {
	result = (*b)[0]
	(*b)[0] = (*b)[len(*b)-1]  // Replace the value at the root with the last leaf
	(*b) = (*b)[:len(*b)-1]    // Shorten slice by one
	b.pushDown(0, LessThan[S]) // Maintain heap property
	return
}

// RemoveMax returns the element with the highest key
// and removes it from the heap.
func (b *MinMaxHeap[S, T]) RemoveMax() (result HeapElement[S, T]) {
	switch len(*b) {
	case 0:
		return
	case 1:
		result = (*b)[0]
		*b = (*b)[:0]
		return
	case 2:
		result = (*b)[1]
		*b = (*b)[:1]
	default:
		i := 1
		if (*b)[1].Key < (*b)[2].Key {
			i = 2
		}
		result = (*b)[i]
		(*b)[i] = (*b)[len(*b)-1]     // Replace the value at the removed node
		*b = (*b)[:len(*b)-1]         // Shorten slice by one
		b.pushDown(i, GreaterThan[S]) // maintain heap property
	}
	return
}

// PeekMin returns the element with the lowest key
// without removing it.
func (b *MinMaxHeap[S, T]) PeekMin() HeapElement[S, T] {
	return (*b)[0]
}

// PeekMax returns the element with the highest key
// without removing it.
func (b MinMaxHeap[S, T]) PeekMax() (result HeapElement[S, T]) {
	switch len(b) {
	case 0:
		return
	case 1:
		result = b[0]
	case 2:
		result = b[1]
	default:
		if b[1].Key < b[2].Key {
			result = b[2]
		} else {
			result = b[1]
		}
	}
	return
}

// Copy returns a shallow copy of the heap
func (b MinMaxHeap[S, T]) Copy() (c MinMaxHeap[S, T]) {
	c = append([]HeapElement[S, T]{}, b...)
	return
}

// GetIterator returns a function that iterates over the elements in the MinMaxHeap b,
// without altering the contents of b.
// If ascending is true the iterator returns the elements in ascending order,
// otherwise the elements are returned in decending order.
// The return value r indicates how many elements remain in the iterator.
// Additional calls to the iterator after all the elements are exhausted (r < 0) will not
// seq fault, but the returned element will contain default values for the Key
// and Data types of the HeapElement.
// Instances of the iterator maintain an internal MaxHeap which will
// gradually grow as more elements are requested. The greatest size of this
// slice occurs near the half way point of the full iteration, and decreases after
// until the size of the slice and the remaining count are zero.
// Depending on the organization of b, the maximum size of the iterator slice map
// be around 1/2 of the length of b, but is usually around 1/3 or less.
func (b MinMaxHeap[S, T]) GetIterator(ascending bool) func() (e HeapElement[S, T], r int) {
	maxHeap := MaxHeap[S, int]{}
	var addToIterator func(i int)
	if ascending {
		addToIterator = func(i int) {
			maxHeap.Insert(HeapElement[S, int]{-b[i].Key, i})
		}
	} else {
		addToIterator = func(i int) {
			maxHeap.Insert(HeapElement[S, int]{b[i].Key, i})
		}
	}
	// Add the initial nodes
	if ascending || len(b) == 1 {
		addToIterator(0)
	} else {
		for i := 1; i < len(b) && i < 3; i++ {
			addToIterator(i)
		}
	}
	remaining := len(b)
	return func() (e HeapElement[S, T], rem int) {
		remaining--
		rem = remaining
		if len(maxHeap) == 0 {
			return
		}
		top := maxHeap.RemoveMax().Data
		e = b[top]
		if IsMinLevel(top) != ascending { // Going back up the tree
			if top > 2 {
				gp := parent(parent(top))
				if leftChild(leftChild(gp)) == top {
					addToIterator(gp)
				}
			}
			return
		}
		l := leftChild(top)
		if l >= len(b) { // top has no children nodes, go back up
			if top > 0 {
				p := parent(top)
				if leftChild(p) == top {
					addToIterator(p)
				}
			}
			return
		}
		// Check both child node's children
		for _, v := range [2]int{rightChild(top), l} {
			vl := leftChild(v)
			if vl < len(b) {
				addToIterator(vl)
				vr := rightChild(v)
				if vr < len(b) {
					addToIterator(vr)
				}
			} else {
				if v < len(b) { // No child nodes go back up
					addToIterator(v)
				}
			}
		}
		return
	}
}

// *****************************************
// Below are some slice helper functions
// *****************************************

// Fill sets all indicies of a to b
func Fill[T any](a []T, b T) {
	for i := range a {
		a[i] = b
	}
}

// Equals returns true if the elements of a and b are the same
func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Reverse reverse of order of s
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
