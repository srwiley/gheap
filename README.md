# gheap
MinHeap, MaxHeap and MinMaxHeap priority queues using golang generics.

This package implements priority queues of the types Min Heap, Max Heap and Min-max Heap (as described by M.D.Atchinson et. al.; https://dl.acm.org/doi/10.1145/6617.6621).
By making use of golang generics, the keys for the queues can be any signed number, integer or float, and the data can be any type. The complete binary trees used by the queues are implicitly encoded in a single slice. Each type implements Insert, RemoveMin and/or RemoveMax, PeekMin and/or PeakMax, and a shallow Copy. The MinMaxHeap type also provides a non-destructive ascending or descenting iterator generating function, GetIterator.

For example, to declare and fill a MinMaxHeap with float32 as Value type and string as Data type:
```go
testSet := []HeapElement[float32, string]{
		{0.0, "spruce"},
		{0.2, "cypress"},
		{0.6, "cedar"},
		{3.0, "oak"},
		{0.14, "maple"},
		{0.14, "palm"},
		{0.6, "mango"},
		{1.0, "walnut"}}

mmheap := MinMaxHeap[float32, string]{}
for _, v := range testSet {
	mmheap.Insert(v)
}

fmt.Println("Max : ", mmheap.PeekMax())
fmt.Println("Min : ", mmheap.PeekMin())
  ```

In addition, generic slice functions Reverse, Equals, and Fill are included.

Please see heap_test.go for more examples.







