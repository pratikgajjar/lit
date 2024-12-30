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

	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	ans := maxScoreWords(words, letters, score)
	Assert(t, ans, 23)

}

func Test2(t *testing.T) {

	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}

	ans := maxScoreWords(words, letters, score)
	Assert(t, ans, 27)

}

func Test3(t *testing.T) {

	words := []string{"leetcode"}
	letters := []byte{'l', 'e', 't', 'c', 'o', 'd'}
	score := []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}

	ans := maxScoreWords(words, letters, score)
	Assert(t, ans, 0)

}
