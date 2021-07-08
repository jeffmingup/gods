package singlylinkedlist

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
)

//单链表实现

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	e := list.first
	for i := 1; i <= index; i, e = i+1, e.next {
	}
	return e.value, true
}
func (list *List) getElement(index int) *element {
	if !list.withinRange(index) {
		return nil
	}
	e := list.first
	for i := 1; i <= index; i++ {
		e = e.next
	}
	return e
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	if index > 0 {
		e := list.getElement(index - 1)
		e.next = e.next.next
		if e.next == nil {
			list.last = e
		}
	} else {
		list.first = list.first.next
	}

	list.size--
}

func (list *List) Add(values ...interface{}) {
	if list.size == 0 {
		e := &element{value: values[0]}
		list.last = e
		list.first = e
		values = values[1:]
		list.size++
	}
	for _, v := range values {
		e := &element{value: v}
		list.last.next = e
		list.last = e
		list.size++
	}

}

func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, v := range values {
		found := false
		for e := list.first; e != nil; e = e.next {
			if v == e.value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	sliceList := list.Values()
	utils.Sort(sliceList, comparator)

	list.Clear()
	list.Add(sliceList)
}

func (list *List) Swap(index1 int, index2 int) {
	if list.withinRange(index1) && list.withinRange(index2) && index1 != index2 {
		var element1, element2 *element
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case index1:
				element1 = currentElement
			case index2:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if index < 0 || index > list.size {
		return
	}
	if index == list.size {
		list.Add(values...)
		return
	}
	var beforeElement, foundElement *element
	if index != 0 {
		beforeElement = list.getElement(index - 1)
		foundElement = beforeElement.next
	} else {
		v := values[0]
		newElement := &element{value: v, next: list.first}
		beforeElement = newElement
		foundElement = list.first
		list.first = newElement
		list.size++
		values = values[1:]
	}

	for _, v := range values {
		newElement := &element{value: v}
		beforeElement.next = newElement
		beforeElement = newElement
	}
	beforeElement.next = foundElement
	list.size += len(values)
}

func (list *List) Set(index int, value interface{}) {
	if index < 0 || index > list.size {
		return
	}
	if index == list.size {
		list.Add(value)
		return
	}
	list.getElement(index).value = value
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) Values() []interface{} {
	if list.size == 0 {
		return []interface{}{}
	}
	values := make([]interface{}, list.size)
	for i, e := 0, list.first; e != nil; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

func (list *List) withinRange(index int) bool {
	return index < list.size && index >= 0
}
