package leet

import (
	"fmt"
	"math"
	"testing"
)

func Assert[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatalf("Failed output=%v, expected=%v", a, b)
		return
	}
	t.Logf("success out=%v", a)
}

// assertFloat performs comparison for floating-point numbers with an epsilon.
func assertFloat(t *testing.T, a, b float64) {
	const epsilon = 1e-3
	if math.Abs(a-b) > epsilon {
		t.Fatalf("Failed output=%v, expected=%v", a, b)
		return
	}
	t.Logf("success out=%v", a)
}

var facts = []Fact{
	{
		src:   "m",
		dest:  "ft",
		coeff: 3.28,
	},
	{
		src:   "ft",
		dest:  "in",
		coeff: 12.0,
	},
	{
		src:   "hr",
		dest:  "min",
		coeff: 60,
	},
	{
		src:   "min",
		dest:  "sec",
		coeff: 60,
	},
}

func Test1(t *testing.T) {

	query := Query{
		value: 2,
		src:   "m",
		dest:  "in",
	}

	ans, e := AnserQuery(query, facts)
	fmt.Println(ans, e)
	Assert(t, ans, 78.72)
}

func Test2(t *testing.T) {

	query := Query{
		value: 13,
		src:   "in",
		dest:  "m",
	}

	ans, e := AnserQuery(query, facts)
	fmt.Println(ans, e)
	assertFloat(t, float64(ans), 0.330)
}

func Test3(t *testing.T) {

	query := Query{
		value: 13,
		src:   "in",
		dest:  "hr",
	}

	ans, e := AnserQuery(query, facts)
	fmt.Println(ans, e)
	if e == nil {
		t.Fatalf("Expected error")
	}
}
