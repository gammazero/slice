package slice

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	c := Copy(s)
	if len(s) != len(c) {
		t.Fatal("copy has wrong lenght:", len(c))
	}

	for i := 0; i < len(s); i++ {
		if s[i] != c[i] {
			t.Fatal("copy is not equal")
		}
	}
}

func TestDelete(t *testing.T) {
	s := []string{"this", "is", "a", "test"}

	s = Delete(s, 2)
	if !reflect.DeepEqual(s, []string{"this", "is", "test"}) {
		t.Fatal("Wrong result for Delete:", s)
	}
}

func TestDeleteFast(t *testing.T) {
	s := []string{"this", "is", "a", "test"}

	s = DeleteFast(s, 1)
	if len(s) != 3 {
		t.Fatal("Wrong number of items")
	}
}

func TestDeleteN(t *testing.T) {
	s := []string{"is", "this", "a", "test", "or", "a", "game"}

	s = DeleteN(s, 3, 3)
	if !reflect.DeepEqual(s, []string{"is", "this", "a", "game"}) {
		t.Fatal("Wrong result for DeleteN:", s)
	}

	s = []string{"is", "this", "a", "test", "or", "a", "game"}
	s = DeleteN(s, 3, 4)
	if !reflect.DeepEqual(s, []string{"is", "this", "a"}) {
		t.Fatal("Wrong result for DeleteN:", s)
	}

	s = DeleteN(s, 1, -1)
	if !reflect.DeepEqual(s, []string{"is"}) {
		t.Fatal("DeleteN should have trucated after first item, got:", s)
	}

	s2 := DeleteN(s, 0, 0)
	if reflect.ValueOf(s).Pointer() != reflect.ValueOf(s2).Pointer() {
		t.Fatal("Should return original slice if deleting 0 items")
	}

	s = []string{"hello", "world"}
	s = DeleteN(s, 2, 5)
	if !reflect.DeepEqual(s, []string{"hello", "world"}) {
		t.Fatal("Wrong result for DeleteN:", s)
	}

	s = DeleteN(s, 0, 10)
	if len(s) != 0 {
		t.Fatal("Expected DeleteN to delete all items")
	}
}

func TestFilter(t *testing.T) {
	s := []string{"this", "is", "a", "test"}

	s = Filter(s, func(word string) bool {
		return len(word) != 0 && word[0] == 't'
	})

	if !reflect.DeepEqual(s, []string{"this", "test"}) {
		t.Fatal("Wrong result for Filter:", s)
	}
}

func TestInsert(t *testing.T) {
	s := []string{"quick", "fox"}
	s = Insert(s, 1, "brown", "furry")
	if !reflect.DeepEqual(s, []string{"quick", "brown", "furry", "fox"}) {
		t.Fatal("Did no get expected result from Insert, got", s)
	}

	s = Insert(s, 0, "the")
	if Index(s, "the") != 0 {
		t.Fatal("Expected \"the\" at index 0")
	}

	assertPanics(t, "Insert", func() {
		Insert(s, -1, "wow")
	})

	assertPanics(t, "Insert", func() {
		s = Insert(s, 6, "wow")
	})

	s = make([]string, 2, 3)
	s[0] = "hello"
	s[1] = "world"
	s2 := Insert(s, 1, "there")
	if reflect.ValueOf(s).Pointer() != reflect.ValueOf(s2).Pointer() {
		t.Fatal("Should not have allocated new slice if original had sufficient capacity")
	}

	s = Insert(nil, 0, "hello")
	if len(s) != 1 {
		t.Fatal("Insert did not insert into nil slice")
	}

	var empty []string
	s = Insert(nil, 0, empty...)
	if s != nil {
		t.Fatal("Insert nothing should return original")
	}
}

func TestPop(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	x, s := Pop(s)
	if x != "baz" {
		t.Fatal("Popped wrong value:", x)
	}
	if len(s) != 2 {
		t.Fatal("Wrong length after Pop:", len(s))
	}
	x, s = Pop(s)
	if x != "bar" {
		t.Fatal("Popped wrong value:", x)
	}
	if len(s) != 1 {
		t.Fatal("Wrong length after Pop:", len(s))
	}
	x, s = Pop(s)
	if x != "foo" {
		t.Fatal("Popped wrong value:", x)
	}
	if len(s) != 0 {
		t.Fatal("Wrong length after Pop:", len(s))
	}
	assertPanics(t, "Pop", func() {
		x, s = Pop(s)
	})
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	Reverse(s)
	if !reflect.DeepEqual(s, []int{7, 6, 5, 4, 3, 2, 1}) {
		t.Fatal("Reverse did not reverse slice:", s)
	}
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		}
	}()

	f()
}
