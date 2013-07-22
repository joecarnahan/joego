package euler

import (
	"container/list"
	"testing"
)

func TestRemoveIfPresent(t *testing.T) {
	myList := list.New()
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 0 {
		t.Errorf("Expected an empty list, got a list of length %v", length)
	}
	myList.PushBack(3)
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 0 {
		t.Errorf("Expected an empty list, got a list of length %v", length)
	}
	myList.PushBack(2)
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 1 {
		t.Errorf("Expected a list of length 1, got a list of length %v", length)
	} else if headElement := myList.Front(); headElement.Value != 2 {
		t.Errorf("Expected an element 2, got %v", headElement.Value)
	}
	myList.PushBack(3)
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 1 {
		t.Errorf("Expected a list of length 1, got a list of length %v", length)
	} else if headElement := myList.Front(); headElement.Value != 2 {
		t.Errorf("Expected an element 2, got %v", headElement.Value)
	}
	myList.PushFront(3)
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 1 {
		t.Errorf("Expected a list of length 1, got a list of length %v", length)
	} else if headElement := myList.Front(); headElement.Value != 2 {
		t.Errorf("Expected an element 2, got %v", headElement.Value)
	}
	myList.PushBack(3)
	myList.PushBack(1)
	removeIfPresent(myList, 3)
	if length := myList.Len(); length != 2 {
		t.Errorf("Expected a list of length 2, got a list of length %v", length)
	} else if headElement := myList.Front(); headElement.Value != 2 {
		t.Errorf("Expected a head element 2, got %v", headElement.Value)
	} else if tailElement := myList.Back(); tailElement.Value != 1 {
		t.Errorf("Expected a tailelement 1, got %v", tailElement.Value)
	}
}

func TestListsEqual(t *testing.T) {
	type ListPair struct {
		first, second *list.List
	}
	empty := list.New()
	singletonA := list.New()
	singletonA.PushBack('a')
	singletonB := list.New()
	singletonB.PushBack('b')
	ab := list.New()
	ab.PushBack('a')
	ab.PushBack('b')
	abCopy := list.New()
	abCopy.PushBack('a')
	abCopy.PushBack('b')
	ba := list.New()
	ba.PushBack('b')
	ba.PushBack('a')
	equalLists := []ListPair{
		ListPair{empty, empty},
		ListPair{singletonA, singletonA},
		ListPair{singletonB, singletonB},
		ListPair{ab, ab},
		ListPair{ab, abCopy},
		ListPair{abCopy, ab},
		ListPair{abCopy, abCopy},
		ListPair{ba, ba}}
	unequalLists := []ListPair{
		ListPair{empty, singletonA},
		ListPair{empty, singletonB},
		ListPair{empty, ab},
		ListPair{empty, abCopy},
		ListPair{empty, ba},
		ListPair{singletonA, empty},
		ListPair{singletonA, singletonB},
		ListPair{singletonA, ab},
		ListPair{singletonA, abCopy},
		ListPair{singletonA, ba},
		ListPair{singletonB, empty},
		ListPair{singletonB, singletonA},
		ListPair{singletonB, ab},
		ListPair{singletonB, abCopy},
		ListPair{singletonB, ba},
		ListPair{ab, empty},
		ListPair{ab, singletonA},
		ListPair{ab, singletonB},
		ListPair{ab, ba},
		ListPair{abCopy, empty},
		ListPair{abCopy, singletonA},
		ListPair{abCopy, singletonB},
		ListPair{abCopy, ba},
		ListPair{ba, empty},
		ListPair{ba, singletonA},
		ListPair{ba, singletonB},
		ListPair{ba, ab},
		ListPair{ba, abCopy}}
	for _, pair := range equalLists {
		if !listsEqual(pair.first, pair.second) {
			t.Errorf("Expected lists %v and %v to be equal", pair.first, pair.second)
		}
	}
	for _, pair := range unequalLists {
		if listsEqual(pair.first, pair.second) {
			t.Errorf("Expected lists %v and %v to be unequal", pair.first, pair.second)
		}
	}
}

func TestAddAll(t *testing.T) {
	empty := list.New()
	emptyCopy := list.New()
	addAll(empty, emptyCopy)
	if length := empty.Len(); length != 0 {
		t.Errorf("Expected result to be empty, got %v", empty)
	}

	singletonA := list.New()
	singletonA.PushBack('a')
	addAll(singletonA, empty)
	if length := empty.Len(); length != 0 {
		t.Errorf("Expected \"empty\" to still be empty, got %v", empty)
	}
	if length := singletonA.Len(); length != 1 {
		t.Errorf("Expected result to contain a single element, got %v", singletonA)
	} else if headElement := singletonA.Front(); headElement.Value != 'a' {
		t.Errorf("Expected single value to be 'a', got '%v'", headElement.Value)
	}

	emptyGetsA := list.New()
	addAll(emptyGetsA, singletonA)
	if length := singletonA.Len(); length != 1 {
		t.Errorf("Expected singletonA to still contain a single element, got %v", singletonA)
	} else if headElement := singletonA.Front(); headElement.Value != 'a' {
		t.Errorf("Expected single value to be 'a', got '%v'", headElement.Value)
	}
	if length := emptyGetsA.Len(); length != 1 {
		t.Errorf("Expected result to contain a single element, got %v", singletonA)
	} else if headElement := emptyGetsA.Front(); headElement.Value != 'a' {
		t.Errorf("Expected single value to be 'a', got '%v'", headElement.Value)
	}

	singletonB := list.New()
	singletonB.PushBack('b')
	aGetsB := list.New()
	aGetsB.PushBack('a')
	ab := list.New()
	ab.PushBack('a')
	ab.PushBack('b')
	addAll(aGetsB, singletonB)
	if !listsEqual(aGetsB, ab) {
		t.Errorf("Expected a + b to be %v, got %v", ab, aGetsB)
	}
}
