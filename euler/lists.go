package euler

import (
	"container/list"
	"fmt"
)

// Stringifies the list. Implemented simply but not efficiently.
func ToString(l *list.List) string {
	result := "List("
	for e := l.Front(); e != nil; e = e.Next() {
		result += fmt.Sprintf("%v", e.Value)
		if e.Next() != nil {
			result += ", "
		}
	}
	result += ")"
	return result
}

// Removes the first occurrence of toRemove from the given list.
func removeIfPresent(fromList *list.List, toRemove interface{}) {
	for e := fromList.Front(); e != nil; e = e.Next() {
		if e.Value == toRemove {
			fromList.Remove(e)
			return
		}
	}
}

// Returns true iff the two given lists contain exactly the same sequence of
// values.
func listsEqual(list1, list2 *list.List) bool {
	if list1.Len() != list2.Len() {
		return false
	}
	e2 := list2.Front();
	for e1 := list1.Front(); e1 != nil; e1 = e1.Next() {
		if e1.Value != e2.Value {
			return false
		}
		e2 = e2.Next()
	}
	return true
}

// Appends all elements of list2 to list1.
func addAll(list1, list2 *list.List) {
	for e := list2.Front(); e != nil; e = e.Next() {
		list1.PushBack(e.Value)
	}
}

// Creates a list that contains all of the values from both lists, where any
// value that appears X times in list1 and Y times in list2 will appear max(X,Y)
// times in the resulting list.
//
// This function modifies the given lists.
// TODO: Determine if it's worth building a defensive copy into this function
func MergeLists(list1, list2 *list.List) *list.List {
	result := list.New()
	for e := list1.Front(); e != nil; e = e.Next() {
		removeIfPresent(list2, e.Value)
		result.PushBack(e.Value)
	}
	addAll(result, list2)
	return result
}
