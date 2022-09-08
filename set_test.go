package zdgoutil

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

	want = []int{0, 5, 5, 8, 9, 1, 4, 7}
	got = append(ExceptLeft(left, right), ExceptRight(left, right)...)

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
