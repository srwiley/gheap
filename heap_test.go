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

func TestMinMaxLoad(t *testing.T) {

	randomElement := func() HeapElement[float64, string] {
		key := rand.NormFloat64()
		return HeapElement[float64, string]{key, "ele: " + fmt.Sprintf("%f", key)}
	}

	mmHeap := MinMaxHeap[float64, string]{}
	for j := 0; j < 5; j++ {
		for i := 0; i < 100; i++ {
			ele := randomElement()
			//ele.Key = float64(j*100 + i)
			mmHeap.Insert(ele)
			if (i+2)%20 == 0 {
				mmHeap.Insert(ele) // add some dups
			}
		}
	}

	iter := mmHeap.GetIterator(false)
	e1, _ := iter()
	lastKey := e1.Key
	for e, r := iter(); r >= 0; e, r = iter() {
		if lastKey < e.Key {
			t.Error("Last key less than current key in decending iterator ", lastKey, e)
		}
	}

	iterA := mmHeap.GetIterator(true)
	a1, _ := iterA()
	lastKey = a1.Key
	for e, r := iterA(); r >= 0; e, r = iterA() {
		if lastKey > e.Key {
			t.Error("Last key greater than current key in acending iterator ", lastKey, e)
		}
	}

	ele1 := mmHeap.RemoveMin()
	lastKey = ele1.Key
	for i := 0; i < 50; i++ {
		peek := mmHeap.PeekMin()
		ele := mmHeap.RemoveMin()
		if peek != ele {
			t.Error("PeekMin and RemoveMin do not match ", peek, ele)
		}
		if lastKey > ele.Key {
			t.Error("Last key greater than current key in RemoveMin ", lastKey, ele)
		}
		lastKey = ele.Key
	}
	ele1 = mmHeap.RemoveMax()
	lastKey = ele1.Key
	for len(mmHeap) > 0 {
		peek := mmHeap.PeekMax()
		ele := mmHeap.RemoveMax()
		if peek != ele {
			t.Error("PeekMax and RemoveMax do not match ", peek, ele)
		}
		if lastKey < ele.Key {
			t.Error("Last key less than current key in RemoveMax ", lastKey, ele, len(mmHeap))
		}
		lastKey = ele.Key
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

	iter := mmheap.GetIterator(true)
	e1, _ := iter()
	lastKey := e1.Key
	for e, r := iter(); r >= 0; e, r = iter() {
		if lastKey > e.Key {
			t.Error("Last key greater than current key in ascending iterator ", lastKey, e)
		}
	}

	iter = mmheap.GetIterator(false)
	a1, _ := iter()
	lastKey = a1.Key
	for e, r := iter(); r >= 0; e, r = iter() {
		if lastKey < e.Key {
			t.Error("Last key less that than current key in decending iterator ", lastKey, e)
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

