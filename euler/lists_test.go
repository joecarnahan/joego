package euler

import (
	"container/list"
	"testing"
)

func TestremoveIfPresent(t *testing.T) {
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

func TestAddAll(t *testing.T) {
	// TODO Implement this
}

func TestMergeLists(t *testing.T) {
	// TODO Implement this
}
