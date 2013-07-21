package euler

import "container/list"

// Removes the first occurrence of toRemove from the given list.
func removeIfPresent(fromList *list.List, toRemove interface{}) {
	for e := fromList.Front(); e != nil; e = e.Next() {
		if e.Value == toRemove {
			fromList.Remove(e)
			return
		}
	}
}

// Appends all elements of list2 to list1.
func AddAll(list1, list2 *list.List) {
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
	AddAll(result, list2)
	return result
}
