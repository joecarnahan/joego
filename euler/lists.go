package euler

import (
	"container/list"
	"fmt"
)

// Stringifies the list. Not a particularly efficient implementation.
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
func RemoveIfPresent(fromList *list.List, toRemove interface{}) {
	for e := fromList.Front(); e != nil; e = e.Next() {
		if e.Value == toRemove {
			fromList.Remove(e)
			return
		}
	}
}

// Returns true iff the two given lists contain exactly the same sequence of
// values.
func ListsEqual(list1, list2 *list.List) bool {
	if list1.Len() != list2.Len() {
		return false
	}
	e2 := list2.Front()
	for e1 := list1.Front(); e1 != nil; e1 = e1.Next() {
		if e1.Value != e2.Value {
			return false
		}
		e2 = e2.Next()
	}
	return true
}

// Appends all elements of list2 to list1.
func AddAll(list1, list2 *list.List) {
	for e := list2.Front(); e != nil; e = e.Next() {
		list1.PushBack(e.Value)
	}
}
