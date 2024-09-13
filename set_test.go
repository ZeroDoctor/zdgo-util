package zdutil

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 7}
	b := []int{0, 2, 3, 5, 5, 8, 9}

	want := []int{2, 3}
	got := Intersect(a, b)

	if len(got) < len(want) {
		fmt.Printf("[ERROR] failed to get correct intersection:\n\t[got=%v]\n\t[want=%v]\n", got, want)
		t.FailNow()
	}

	for i := range want {
		if want[i] != got[i] {
			fmt.Printf("[ERROR] failed to get correct intersection:\n\t[got=%v]\n\t[want=%v]\n", got, want)
			t.FailNow()
			break
		}
	}

	seta := NewSet(a...)
	setb := NewSet(b...)

	setgot := seta.Intersect(*setb)

	if setgot.Len() < len(want) {
		fmt.Printf("[ERROR] failed to get correct length intersection:\n\t[got=%v]\n\t[want=%v]\n", setgot.Values(), want)
		t.FailNow()
	}

	if !setgot.Contains(want...) {
		fmt.Printf("[ERROR] failed to get correct intersection:\n\t[got=%v]\n\t[want=%v]\n", setgot.Values(), want)
		t.FailNow()
	}
}

func TestExcept(t *testing.T) {
	left := []int{1, 2, 3, 3, 4, 7}
	right := []int{0, 2, 3, 5, 5, 8, 9}

	want := []int{1, 4, 7}
	got := ExceptRight(left, right)

	if len(got) < len(want) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
		t.FailNow()
	}

	for i := range want {
		if want[i] != got[i] {
			fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
			t.FailNow()
			break
		}
	}

	leftset := NewSet(left...)
	rightset := NewSet(right...)

	want = []int{1, 4, 7}
	gotset := rightset.ExceptRight(*leftset)

	if gotset.Len() < len(want) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", gotset.Values(), want)
		t.FailNow()
	}

	if !gotset.Contains(want...) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", gotset.Values(), want)
		t.FailNow()
	}

	want = []int{0, 5, 5, 8, 9}
	got = ExceptLeft(left, right)

	if len(got) < len(want) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
		t.FailNow()
	}

	for i := range want {
		if want[i] != got[i] {
			fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
			t.FailNow()
			break
		}
	}

	want = []int{0, 5, 8, 9}
	gotset = rightset.ExceptLeft(*leftset)

	if gotset.Len() < len(want) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", gotset.Values(), want)
		t.FailNow()
	}

	if !gotset.Contains(want...) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", gotset.Values(), want)
		t.FailNow()
	}

	want = []int{0, 5, 5, 8, 9, 1, 4, 7}
	got = Union(ExceptLeft(left, right), ExceptRight(left, right))

	if len(got) < len(want) {
		fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
		t.FailNow()
	}

	for i := range want {
		if want[i] != got[i] {
			fmt.Printf("[ERROR] failed to get correct exception:\n\t[got=%v]\n\t[want=%v]\n", got, want)
			t.FailNow()
			break
		}
	}
}
