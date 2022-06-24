package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Count counts the number of instances of item in s.
func Count[T constraints.Ordered](s []T, item T) int {
	var count int
	for _, elem := range s {
		if elem == item {
			count++
		}
	}
	return count
}

// Cut slices s around the first instance of sep, returning the slice before
// and after sep. The found result reports whether sep is an element of s. If
// sep is not an element of s, cut returns s, nil, false.
func Cut[T constraints.Ordered](s []T, sep T) (before, after []T, found bool) {
	if i := Index(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, nil, false
}

// Remove returns a slize with all occurrances of each item removed.
func Remove[T constraints.Ordered](s []T, items ...T) []T {
	if len(s) == 0 {
		return s
	}
	var n int
	if len(items) == 1 {
		item := items[0]
		for i := range s {
			if s[i] != item {
				s[n] = s[i]
				n++
			}
		}
	} else if len(items) > 16 {
		// If there are many items to remove, it is more efficient to look them
		// up in a map.
		set := make(map[T]struct{}, len(s))
		for i := range items {
			set[items[i]] = struct{}{}
		}
		for i := range s {
			if _, ok := set[s[i]]; !ok {
				s[n] = s[i]
				n++
			}
		}
	} else if len(items) == 0 {
		return s
	} else {
		for i := range s {
			found := false
			for j := range items {
				if s[i] == items[j] {
					found = true
					break
				}
			}
			if !found {
				s[n] = s[i]
				n++
			}
		}
	}

	// Remove unused items that are still referenced by slice memory.
	var zero T
	for i := n; i < len(s); i++ {
		s[i] = zero
	}
	return s[:n]
}

// Index returns the index of the first instance of item in s, or -1 if item is
// not present in s.
func Index[T constraints.Ordered](s []T, item T) int {
	for i := 0; i < len(s); i++ {
		if s[i] == item {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of item in s, or -1 if item is
// not present in s.
func LastIndex[T constraints.Ordered](s []T, item T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == item {
			return i
		}
	}
	return -1
}

// Replace replaces the first n instances of old, in s, with new. If n is -1,
// then there is no limit on the number of replacements.
func Replace[T constraints.Ordered](s []T, old, new T, n int) {
	if n != 0 && old != new {
		for i := 0; i < len(s); i++ {
			if s[i] == old {
				s[i] = new
				n--
				if n == 0 {
					break
				}
			}
		}
	}
}

// ReplaceLast replaces the last n instances of old, in s, with new. If n is -1,
// then there is no limit on the number of replacements.
func ReplaceLast[T constraints.Ordered](s []T, old, new T, n int) {
	if n != 0 && old != new {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == old {
				s[i] = new
				n--
				if n == 0 {
					break
				}
			}
		}
	}
}

// Sort sorts s in-place in ascending order.
func Sort[T constraints.Ordered](s []T) {
	srt := sorter[T]{
		items: s,
	}
	sort.Sort(srt)
}

// SortReverse sorts s in-place in descending order.
func SortReverse[T constraints.Ordered](s []T) {
	srt := sorter[T]{
		items: s,
	}
	sort.Sort(sort.Reverse(srt))
}

// Unique removes duplicate items in s, keeping only the first instance of
// each.
func Unique[T constraints.Ordered](s []T) []T {
	if len(s) > 1 {
		seen := make(map[T]struct{}, len(s))
		n := 0
		for i := range s {
			if _, ok := seen[s[i]]; !ok {
				seen[s[i]] = struct{}{}
				s[n] = s[i]
				n++
			}
		}
		var zero T
		for i := n; i < len(s); i++ {
			s[i] = zero
		}
		s = s[:n]
	}
	return s
}

type sorter[T constraints.Ordered] struct {
	items []T
}

func (s sorter[T]) Len() int           { return len(s.items) }
func (s sorter[T]) Less(i, j int) bool { return s.items[i] < s.items[j] }
func (s sorter[T]) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }
