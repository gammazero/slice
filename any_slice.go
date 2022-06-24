package slice

// Copy makes a copy of a slice.
func Copy[T any](s []T) []T {
	cpy := make([]T, len(s))
	copy(cpy, s)
	return cpy
}

// Delete removes the item at index i.
func Delete[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

// DeleteFast removes the item at index i without preserving order.
func DeleteFast[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

// DeleteN removes items in s from index start to start + n. If n is -1, then s
// is truncated at start.
func DeleteN[T any](s []T, start, n int) []T {
	if n == 0 {
		return s
	}
	var end int
	if n == -1 {
		end = len(s)
	} else {
		end = start + n
		if end > len(s) {
			end = len(s)
		}
	}
	copy(s[start:], s[end:])
	// Zero cut values still referenced by slice to prevent memory leaks.
	var zero T
	for i, n := len(s)-end+start, len(s); i < n; i++ {
		s[i] = zero
	}
	return s[:len(s)-end+start]
}

// Filter removes all items for which keep returns false.
func Filter[T any](s []T, keep func(T) bool) []T {
	n := 0
	for i := range s {
		if keep(s[i]) {
			s[n] = s[i]
			n++
		}
	}
	var zero T
	for i := n; i < len(s); i++ {
		s[i] = zero
	}
	return s[:n]
}

// Insert inserts items into s at the index i.
func Insert[T any](s []T, i int, items ...T) []T {
	if len(items) == 0 {
		return s
	}
	if n := len(s) + len(items); n <= cap(s) {
		s2 := s[:n]
		// Copy rhs past gap of size n.
		copy(s2[i+len(items):], s[i:])
		// Copy items into gap.
		copy(s2[i:], items)
		return s2
	}
	s2 := make([]T, len(s)+len(items))
	copy(s2, s[:i])
	copy(s2[i:], items)
	copy(s2[i+len(items):], s[i:])
	return s2
}

// Pop removes and returns the last item in the slice.
func Pop[T any](s []T) (T, []T) {
	x := s[len(s)-1]
	var zero T
	s[len(s)-1] = zero
	return x, s[:len(s)-1]
}

// Reverse reverses the order of items in s, in place.
func Reverse[T any](s []T) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
