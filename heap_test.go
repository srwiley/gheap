package gheap_test

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/srwiley/gheap"
)

func TestLevelfunc(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(i, " ", IsMinLevel(i))
	}
}

func writeHeapF(best MinMaxHeap[float64, string]) {
	iter2 := best.GetIterator(false)
	for ef, r2 := iter2(); r2 >= 0; ef, r2 = iter2() {
		fmt.Println("elt", ef, r2)
	}
}

func TestMinMaxLoad(t *testing.T) {
	randomElement := func() HeapElement[float64, string] {
		key := rand.NormFloat64()
		return HeapElement[float64, string]{key, "ele: " + fmt.Sprintf("%f", key)}
	}

	mmHeap := MinMaxHeap[float64, string]{}
	for i := 0; i < 2; i++ {
		ele := randomElement()
		mmHeap.Insert(ele)
	}
	for j := 0; j < 1; j++ {
		for i := 0; i < 200; i++ {
			ele := randomElement()
			mmHeap.Insert(ele)
			if (i+2)%20 == 0 {
				mmHeap.Insert(ele) // add some dups
			}
			iter := mmHeap.GetIterator(false)
			e1, _ := iter()
			lastKey := e1.Key
			eCount := 1
			for e, r := iter(); r >= 0; e, r = iter() {
				eCount++
				if lastKey < e.Key {
					t.Error("Last key less than current key in decending iterator ", lastKey, e.Key, "len", len(mmHeap))
					writeHeapF(mmHeap)
					return

				}
				lastKey = e.Key
			}
			if eCount != len(mmHeap) {
				t.Error("iteration count and heap length not equal", eCount, len(mmHeap))
			}
		}
	}

	for i := 0; i < 30; i++ {
		iter := mmHeap.GetIterator(false)
		e1, _ := iter()
		lastKey := e1.Key
		eCount := 1
		for e, r := iter(); r >= 0; e, r = iter() {
			eCount++
			if lastKey < e.Key {
				t.Error("Last key less than current in decending iterator ", lastKey, e.Key, "len", len(mmHeap))
				writeHeapArrayf(mmHeap)
				return

			}
			lastKey = e.Key
		}
		if eCount != len(mmHeap) {
			t.Error("it count and heap length not equal", eCount, len(mmHeap))
		}

		if i%2 == 0 {
			peek := mmHeap.PeekMin()
			ele := mmHeap.RemoveMin()
			//fmt.Println("removed min", ele)
			if peek != ele {
				t.Error("PeekMin and RemoveMin do not match ", peek, ele)
			}
			continue
		}
		peek := mmHeap.PeekMax()
		ele := mmHeap.RemoveMax()
		if peek != ele {
			t.Error("PeekMax and RemoveMax do not match ", peek, ele)
		}

	}
}

func TestMinMaxHeap(t *testing.T) {
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

	iter := mmheap.GetIterator(true)
	e1, _ := iter()
	lastKey := e1.Key
	for e, r := iter(); r >= 0; e, r = iter() {
		if lastKey > e.Key {
			t.Error("Last key greater than current key in ascending iterator ", lastKey, e)
		}
		lastKey = e.Key
	}

	iter = mmheap.GetIterator(false)
	a1, _ := iter()
	lastKey = a1.Key
	for e, r := iter(); r >= 0; e, r = iter() {
		if lastKey < e.Key {
			t.Error("Last key less that than current key in decending iterator ", lastKey, e)
		}
		lastKey = e.Key
	}
}

func writeHeap(best MinMaxHeap[int, string]) {
	iter2 := best.GetIterator(false)
	for ef, r2 := iter2(); r2 >= 0; ef, r2 = iter2() {
		fmt.Println("elt", ef)
	}

}

func writeHeapArrayf(best MinMaxHeap[float64, string]) {
	for i, v := range best {
		fmt.Println(i, ":", v)
	}
}

func TestMinMaxHeapI(t *testing.T) {
	testSet := []HeapElement[int, string]{
		{0, "a"},
		{2, "b"},
		{1, "c"},
		{1, "d"},
		{1, "e"},
		{3, "f"},
		{1, "g"},
		{2, "h"},
		{1, "i"},
		{1, "j"},
		{2, "k"},
		{1, "l"},
		{2, "m"},
		{2, "n"},
		{2, "o"},
		{3, "p"},
	}

	mmheap := MinMaxHeap[int, string]{}
	for _, v := range testSet {
		mmheap.Insert(v)
		fmt.Println()
		writeHeap(mmheap)
		if len(mmheap) > 6 {
			mmheap.RemoveMin()
		}
		iter := mmheap.GetIterator(false)
		a1, _ := iter()
		lastKey := a1.Key
		toErr := false
		for e, r := iter(); r >= 0; e, r = iter() {
			if lastKey < e.Key {
				t.Error("Last key less that than current key in decending iterator ", lastKey, e)
				toErr = true
			}
			lastKey = e.Key
		}
		if toErr {
			return
		}
	}
}

func TestMaxHeap(t *testing.T) {
	testSet := []HeapElement[float32, string]{
		{0.0, "spruce"},
		{0.2, "cypress"},
		{0.6, "cedar"},
		{3.0, "oak"},
		{0.14, "maple"},
		{1.0, "walnut"}}

	maxheap := MaxHeap[float32, string]{}
	for _, v := range testSet {
		maxheap.Insert(v)
		//maxheap,(v)
	}
	for len(maxheap) > 0 {
		ele := maxheap.RemoveMax()
		fmt.Println(ele)
	}
	fmt.Println("reversing keys")
	for _, v := range testSet {
		v.Key = -v.Key
		maxheap.Insert(v)
	}
	lastKey := maxheap.PeekMax().Key
	for len(maxheap) > 0 {
		peek := maxheap.PeekMax()
		ele := maxheap.RemoveMax()
		if peek != ele {
			t.Error("PeekMax and RemoveMax do not match ", peek, ele)
		}
		if lastKey < ele.Key {
			t.Error("lastKey < curent during maxHeap remove ", peek, ele)
		}
		lastKey = ele.Key
		fmt.Println(ele)
	}
}

func TestMinHeap(t *testing.T) {
	testSet := []HeapElement[float32, string]{
		{0, "spruce"},
		{2, "cypress"},
		{6, "cedar"},
		{30, "oak"},
		{1, "maple"},
		{10, "walnut"}}

	minheap := MinHeap[float32, string]{}
	for _, v := range testSet {
		minheap.Insert(v)
	}
	for len(minheap) > 0 {
		ele := minheap.RemoveMin()
		fmt.Println(ele)
	}
	fmt.Println("reversing keys")
	for _, v := range testSet {
		v.Key = -v.Key
		minheap.Insert(v)
	}
	lastKey := minheap.PeekMin().Key
	for len(minheap) > 0 {
		peek := minheap.PeekMin()
		ele := minheap.RemoveMin()
		if peek != ele {
			t.Error("PeekMax and RemoveMax do not match ", peek, ele)
		}
		if lastKey > ele.Key {
			t.Error("lastKey < curent during maxHeap remove ", peek, ele)
		}
		lastKey = ele.Key
		fmt.Println(ele)
	}
}

func TestHelpersHeap(t *testing.T) {

	testSet := []HeapElement[float32, string]{
		{0, "spruce"},
		{2, "cypress"},
		{6, "cedar"},
		{30, "oak"},
		{1, "maple"},
		{10, "walnut"}}

	testSetRev := []HeapElement[float32, string]{
		{10, "walnut"},
		{1, "maple"},
		{30, "oak"},
		{6, "cedar"},
		{2, "cypress"},
		{0, "spruce"}}

	baseElement := HeapElement[float32, string]{float32(0), "empty"}

	Reverse(testSet)
	if !Equals(testSet, testSetRev) {
		t.Error("reverse slice failed ", testSet, " vs ", testSetRev)
	}
	Reverse(testSet)
	if Equals(testSet, testSetRev) {
		t.Error("double reverse slice failed ", testSet, " vs ", testSetRev)
	}

	Fill(testSet, baseElement)
	for _, e := range testSet {
		if e != baseElement {
			t.Error("Set to failed ", e, " vs ", baseElement)
		}
	}

}
