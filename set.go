package zdutil

type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates a new set containing the specified elements.
// The function initializes a set by adding each of the provided elements.
// Duplicates in the input are ignored, and only distinct elements are stored.
func NewSet[T comparable](elems ...T) *Set[T] {
	elements := make(map[T]struct{}, len(elems))

	for i := range elems {
		elements[elems[i]] = struct{}{}
	}

	return &Set[T]{
		elements: elements,
	}
}

// Add adds the specified elements to the set.
//
// The function is a no-op if any of the specified elements already exist in the set.
func (s *Set[T]) Add(elems ...T) {
	for i := range elems {
		s.elements[elems[i]] = struct{}{}
	}
}

// Remove deletes the specified element from the set.
// If the element does not exist in the set, the function does nothing.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if the set contains all of the given elements.
//
// If len(elems) is 0 and the set is not empty, the function returns false.
// Otherwise, the function returns true if the set contains all of the given elements.
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

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.elements)
}

// Clear resets the set to its initial state, removing all elements.
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// Values returns all elements in the set.
func (s *Set[T]) Values() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

// Intersect returns a new set that is the intersection of the two sets.
// The function creates a new set and adds all elements from the two sets
// to the new set. The function ignores duplicates.
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

// Union returns a new set that is the union of the two sets. The function
// creates a new set and adds all elements from both sets to the new set.
// The function ignores duplicates.
func (s *Set[T]) Union(a Set[T]) Set[T] {
	union := NewSet[T]()

	union.Add(a.Values()...)

	return *union
}

// ExceptRight returns elements of the first set that do not exist in the second set.
// Ignores duplicates.
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

// ExceptLeft returns elements of the first set that do not exist in the second set.
// Ignores duplicates.
func (s *Set[T]) ExceptLeft(a Set[T]) Set[T] {
	return a.ExceptRight(*s)
}

// Intersect returns a slice containing elements that are present in both input slices a and b.
// The function maintains the order of elements as they appear in slice b and does not include duplicates.

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

// Union returns a slice containing all elements from the input slices a and b.
// The function includes duplicates.
func Union[T comparable](a, b []T) []T {
	return append(a, b...)
}

// ExceptRight returns elements of the first array that do not exist in the second array.
// Ignores duplicates and the order of elements in the second array.

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

// ExceptLeft returns elements of a that do not exist in b. Ignores duplicates.
func ExceptLeft[T comparable](a, b []T) []T {
	return ExceptRight(b, a)
}
