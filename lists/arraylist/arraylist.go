package arraylist

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
)

type List struct {
	elements []interface{}
	size     int
}

const growthFactor = float32(2.0)
const shrinkFactor = float32(0.25)

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List) Contains(values ...interface{}) bool {
	for _, searchValue := range values {
		found := false
		for _, e := range list.elements {
			if e == searchValue {
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
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

func (list *List) Swap(index1, index2 int) {
	if list.withinRange(index1) && list.withinRange(index2) {
		list.elements[index1], list.elements[index2] = list.elements[index2], list.elements[index1]
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	l := len(values)
	list.growBy(len(values))
	copy(list.elements[index+l:], list.elements[index:list.size])
	copy(list.elements[index:], values)
	list.size += l
}

//超过索引大小，则不做处理
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return int(list.size)
}

func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) Values() []interface{} {
	return list.elements[:list.size]
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, v := range values {
		list.elements[list.size] = v
		list.size++
	}
}

//数组扩容
func (list *List) growBy(n int) {
	// if cap(list.elements)+n <= list.Size() {
	// 	return
	// }
	// //小心list.size的初始零值，可以直接加上size
	// newCapacity := (list.size + n) * int(GrowthFactor)
	// list.reSize(newCapacity)
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.reSize(newCapacity)
	}
}

//数组缩容
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.reSize(list.size)
	}
}

//数组地址重新分配
func (list *List) reSize(cap int) {
	newElements := make([]interface{}, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}
func (list *List) withinRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for k, v := range list.elements {
		if v == value {
			return k
		}
	}
	return -1
}
