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
	arr := []int{1, 3, 5}
	ans := findMin(arr)
	if ans != 1 {
		t.Fatalf("Failed ans=%v", ans)
	}
}

func Test2(t *testing.T) {
	arr := []int{2, 2, 2, 0, 1}
	ans := findMin(arr)
	if ans != 0 {
		t.Fatalf("Failed ans=%v", ans)
	}
}
