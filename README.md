# gheap
MinHeap, MaxHeap and MinMaxHeap priority queues using golang generics.

This package implements priority queues of the types Min Heap, Max Heap and Min-max Heap (as described by M.D.Atchinson; https://dl.acm.org/doi/10.1145/6617.6621).
By making use of goland generics, the keys for the queues can be any signed number, integer or float, and the data can be any type, so this implementation can used for a variety of applications. The complete binary trees used by the queues are implicitly encoded in a golang slice. Each type implements Insert, RemoveMin and/or RemoveMax, and PeekMin and/or PeakMax, and a shallow Copy. The MinMaxHeap type also provides a non-destructive ascending or descenting iterator generating function.

In addition, generic slice functions Reverse, Equals, and Fill are included.

Please see heap_test.go for examples.







