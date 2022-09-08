package zdutil

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
func Union[T comparable](a, b []T) {}

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
	var except []T

	commonMap := make(map[T]bool)

	for i := range a {
		commonMap[a[i]] = true
	}

	for i := range b {
		if !commonMap[b[i]] {
			except = append(except, b[i])
		}
	}

	return except
}
