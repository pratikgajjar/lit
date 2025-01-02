package main

import (
	"container/heap"
)

type heapInt []int

func (h heapInt) Len() int           { return len(h) }
func (h heapInt) Less(i, j int) bool { return h[i] > h[j] }
func (h heapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapInt) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *heapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// func (h heapInt) Len() int { return len(h) }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	h := &heapInt{}
	heap.Init(h)

	i, j := 0, 0

	total := len(nums1) + len(nums2)
	for h.Len() < (total+2)/2 {
		if i < len(nums1) && (j >= len(nums2) || len(nums2) == 0 || nums1[i] <= nums2[j]) {
			h.Push(nums1[i])
			i += 1
			continue
		}
		if j < len(nums2) && (i >= len(nums1) || len(nums1) == 0 || nums1[i] > nums2[j]) {
			h.Push(nums2[j])
			j += 1
		}
	}

	if h.Len() == 0 {
		return 0
	}
	if h.Len() == 1 {
		return float64(h.Pop().(int))
	}

	i, j = h.Pop().(int), h.Pop().(int)
	if total%2 == 0 {
		return float64(i+j) / 2
	} else {
		return float64(i)
	}
}
