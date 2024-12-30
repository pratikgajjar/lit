package main

import (
	"fmt"
)

func dp(words []string, budget map[rune]int, score []int, ch int) int {
	if ch > len(words)-1 {
		return 0
	}
	cost := 0
	w := words[ch]
	reset := map[rune]int{}
	// either we choose whole word or skip with score reset
	for _, l := range w {
		if budget[l] > 0 {
			cost += score[l-'a']
			reset[l] += 1
			budget[l] -= 1
		} else {
			cost = 0
			break
		}
	}
	for r := range budget {
		reset[r] += budget[r]
	}
	// fmt.Println(reset, "reset")
	// fmt.Println(budget, "budget")
	// return 0
	return max(dp(words, reset, score, ch+1), cost+dp(words, budget, score, ch+1))
}

func possible(word string, budget map[rune]int) bool {
	req := map[rune]int{}
	for _, w := range word {
		req[w] += 1
	}
	for r := range req {
		if req[r] > budget[r] {
			return false
		}
	}
	return true
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	lf := map[rune]int{}
	for _, l := range letters {
		lf[rune(l)] += 1
	}

	fmt.Println(lf, "budget - origin")
	goodWords := []string{}

	for _, w := range words {
		if possible(w, lf) {
			goodWords = append(goodWords, w)
		}
	}

	return dp(goodWords, lf, score, 0)

}
