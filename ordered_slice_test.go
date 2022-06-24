package slice

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	before, after, found := Cut(s, "jumps")
	if !found {
		t.Fatal("expected true")
	}

	if !reflect.DeepEqual(before, []string{"the", "quick", "brown", "fox"}) {
		t.Fatal("Did not get expected before, got", before)
	}

	if !reflect.DeepEqual(after, []string{"over", "the", "lazy", "dog"}) {
		t.Fatal("Did not get expected after, got", after)
	}

	before, after, found = Cut(s, "xyz")
	if found {
		t.Fatal("expected false")
	}
	if !reflect.DeepEqual(before, s) {
		t.Fatal("Did not get expected before, got", before)
	}
	if after != nil {
		t.Fatal("Expected nil")
	}
}

func TestIndex(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	if Index(s, "the") != 0 {
		t.Fatal("wrong value from Index")
	}
	if RIndex(s, "the") != 6 {
		t.Fatal("wrong value from RIndex")
	}
	if Index(s, "xyz") != -1 {
		t.Fatal("wrong value from Index")
	}
	if RIndex(s, "xyz") != -1 {
		t.Fatal("wrong value from RIndex")
	}
}

func TestCount(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	if Count(s, "the") != 2 {
		t.Fatal("Wrong value from Count")
	}
	if Count(s, "fox") != 1 {
		t.Fatal("Wrong value from Count")
	}
	if Count(s, "xyz") != 0 {
		t.Fatal("Wrong value from Count")
	}
}

func TestRemove(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	s = Remove(s, "the", "brown", "lazy")
	if !reflect.DeepEqual(s, []string{"quick", "fox", "jumps", "over", "dog"}) {
		t.Fatal("Wrong result after remove:", s)
	}
	s2 := Remove(s)
	if reflect.ValueOf(s).Pointer() != reflect.ValueOf(s2).Pointer() {
		t.Fatal("Should return original slice if removing no items")
	}
	s2 = Remove(s, "nothere")
	if reflect.ValueOf(s).Pointer() != reflect.ValueOf(s2).Pointer() {
		t.Fatal("Should return original slice if no items removed")
	}

	a := make([]int, 64)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	rm := make([]int, 63)
	copy(rm, a[1:])

	a = Remove(a, rm...)
	if len(a) != 1 {
		t.Fatal("Expected one value to remain in s, got:", a)
	}
	a = Remove(a, 0)
	if len(a) != 0 {
		t.Fatal("Expected empty slice, got:", a)
	}

	a2 := Remove(a, 0)
	if reflect.ValueOf(a).Pointer() != reflect.ValueOf(a2).Pointer() {
		t.Fatal("Removing from zero-length slice should return slice")
	}
}

func TestReplace(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	Replace(s, "the", "a", -1)
	if !reflect.DeepEqual(s, []string{"a", "quick", "brown", "fox", "jumps", "over", "a", "lazy", "dog"}) {
		t.Fatal("Did not get expected after, got", s)
	}

	Replace(s, "a", "my", 1)
	if !reflect.DeepEqual(s, []string{"my", "quick", "brown", "fox", "jumps", "over", "a", "lazy", "dog"}) {
		t.Fatal("Did not get expected after, got", s)
	}

	Replace(s, "my", "a", 1)
	RReplace(s, "a", "your", 1)
	if !reflect.DeepEqual(s, []string{"a", "quick", "brown", "fox", "jumps", "over", "your", "lazy", "dog"}) {
		t.Fatal("Did not get expected after, got", s)
	}

	s = []string{"foo", "bar", "bar", "bar", "baz", "baz", "baz"}
	Replace(s, "bar", "aaa", 2)
	RReplace(s, "baz", "zzz", 2)
	if !reflect.DeepEqual(s, []string{"foo", "aaa", "aaa", "bar", "baz", "zzz", "zzz"}) {
		t.Fatal("Did not get expected result, got", s)
	}
}

func TestSort(t *testing.T) {
	s := []int{6, 2, 1, 7, 3, 5, 4}
	Sort(s)
	if !reflect.DeepEqual(s, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Fatal("Not sorted in ascending order:", s)
	}
	RSort(s)
	if !reflect.DeepEqual(s, []int{7, 6, 5, 4, 3, 2, 1}) {
		t.Fatal("Not sorted in descending  order:", s)
	}
}

func TestUnique(t *testing.T) {
	s := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	before := len(s)

	s = Unique(s)
	if Count(s, "the") != 1 {
		t.Fatal("Unique did not remove duplicate")
	}

	if len(s) != before-1 {
		t.Fatal("Unique should only have removed 1 duplicate, result:", s)
	}

	prevIndex := -1
	for _, word := range []string{"the", "quick", "brown", "fox", "jumps", "over", "lazy", "dog"} {
		index := Index(s, word)
		if index <= prevIndex {
			t.Fatal("Order not preserved by UniqueStable:", s)
		}
		prevIndex = index
	}
}

func TestUniqueInt(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 1, 7, 8}

	before := len(s)

	s = Unique(s)
	if Count(s, 1) != 1 {
		t.Fatal("Unique did not remove duplicate")
	}

	if len(s) != before-1 {
		t.Fatal("Unique should only have removed 1 duplicate, result:", s)
	}

	prevIndex := -1
	for i := 1; i < 9; i++ {
		index := Index(s, i)
		if index <= prevIndex {
			t.Fatal("Order not preserved by UniqueStable:", s)
		}
		prevIndex = index
	}
}

func BenchmarkRemove(b *testing.B) {
	// Start workers, and have them all wait on a channel before completing.
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int, 1024)
		for i := 0; i < len(s); i++ {
			s[i] = i
		}
		//rm := Copy(s)
		b.StartTimer()
		s = Remove(s, 10, 20, 30, 100, 200, 300, 400, 500, 111, 222, 333, 444, 555, 666, 777, 888, 999, 1000, 1100, 1200, 1300, 1400)
		if len(s) == 0 {
			panic("empty result")
		}
	}
}
