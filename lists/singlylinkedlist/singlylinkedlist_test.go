package singlylinkedlist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/emirpasic/gods/utils"
)

func TestNew(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		// TODO: Add test cases.
		{name: "1", args: args{values: []interface{}{"a", 1, "c"}}, want: &List{
			first: &element{
				value: "a",
				next:  &element{value: 1, next: &element{value: "c", next: nil}},
			},
			last: &element{
				value: "c",
				next:  nil,
			},
			size: 3,
		}},
		{name: "2", args: args{values: []interface{}{"v", "2323"}}, want: &List{
			first: &element{
				value: "v",
				next:  &element{value: "2323", next: nil},
			},
			last: &element{
				value: "2323",
				next:  nil,
			},
			size: 2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestListAdd(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListRemove(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListSwap(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	list.Swap(0, 1)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListSort(t *testing.T) {
	list := New()
	list.Sort(utils.StringComparator)
	list.Add("e", "f", "g", "a", "b", "c", "d")
	list.Sort(utils.StringComparator)
	for i := 1; i < list.Size(); i++ {
		a, _ := list.Get(i - 1)
		b, _ := list.Get(i)
		if a.(string) > b.(string) {
			t.Errorf("Not sorted! %s > %s", a, b)
		}
	}
}

func TestListClear(t *testing.T) {
	list := New()
	list.Add("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListContains(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	list.Clear()
	if actualValue := list.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Contains("a", "b", "c"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestListValues(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.Values()...), "abc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListInsert(t *testing.T) {
	list := New()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", list.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListSet(t *testing.T) {
	list := New()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.Values()...), "abbc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

func BenchmarkSinglyLinkedListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}
