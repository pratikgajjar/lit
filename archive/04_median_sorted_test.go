package main

import (
	"testing"
)

func Assert[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatalf("Failed output=%v, expected=%v", a, b)
		return
	}
}

func Test1(t *testing.T) {
	a1 := []int{1, 3}
	a2 := []int{2}
	ans := findMedianSortedArrays(a1, a2)
	t.Fatalf("failed %v", ans)
}

func Test2(t *testing.T) {
	a1 := []int{1, 3}
	a2 := []int{2, 7}
	ans := findMedianSortedArrays(a1, a2)
	t.Fatalf("failed %v", ans)
}
