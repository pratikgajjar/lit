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

}
