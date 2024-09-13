package zdutil

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](elems ...T) *Set[T] {
	elements := make(map[T]struct{}, len(elems))

	for i := range elems {
		elements[elems[i]] = struct{}{}
	}

	return &Set[T]{
		elements: elements,
	}
}

func (s *Set[T]) Add(elems ...T) {
	for i := range elems {
		s.elements[elems[i]] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(elems ...T) bool {
	if len(elems) == 0 && s.Len() > 0 {
		return false
	}

	i := 0
	exists := true
	for exists && i < len(elems) {
		_, exists = s.elements[elems[i]]
		i++
	}

	return exists
}

func (s *Set[T]) Len() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) Values() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set[T]) Intersect(a Set[T]) Set[T] {
	intersect := NewSet[T]()

	values := a.Values()
	for i := range values {
		if s.Contains(values[i]) {
			intersect.Add(values[i])
		}
	}

	return *intersect
}

func (s *Set[T]) Union(a Set[T]) Set[T] {
	union := NewSet[T]()

	union.Add(a.Values()...)

	return *union
}

func (s *Set[T]) ExceptRight(a Set[T]) Set[T] {
	except := NewSet[T]()

	values := a.Values()
	for i := range values {
		if !s.Contains(values[i]) {
			except.Add(values[i])
		}
	}

	return *except
}

func (s *Set[T]) ExceptLeft(a Set[T]) Set[T] {
	return a.ExceptRight(*s)
}

// Intersect returns an array of elements that are common between two arrays.
// Ignores duplicates
func Intersect[T comparable](a, b []T) []T {
	var intersect []T

	commonMap := make(map[T]bool)

	for i := range a {
		commonMap[a[i]] = true
	}

	for i := range b {
		if commonMap[b[i]] {
			intersect = append(intersect, b[i])
		}
	}

	return intersect
}

// Union just use append() like a normal person
func Union[T comparable](a, b []T) []T { return append(a, b...) }

// ExceptRight returns an array of elements only in the first array
// Doesn't ignore duplicates
func ExceptRight[T comparable](a, b []T) []T {
	var except []T

	commonMap := make(map[T]bool)

	for i := range b {
		commonMap[b[i]] = true
	}

	for i := range a {
		if !commonMap[a[i]] {
			except = append(except, a[i])
		}
	}

	return except
}

// ExceptRight returns an array of elements only in the second array.
// More of a convenience function. Could just use ExceptRight and switch the params
// Doesn't ignore duplicates
func ExceptLeft[T comparable](a, b []T) []T {
	return ExceptRight(b, a)
}
