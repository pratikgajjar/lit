package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(k int) int {
	if k > 0 {
		return k
	} else {
		return -k
	}
}

func main() {
	left := []int{}
	right := map[int]int{}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Stop reading on EOF
		}
		nums := strings.Fields(line)
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right[r] += 1
	}

	ans := 0

	for _, l := range left {
		freq, _ := right[l]
		ans += (freq * l)
	}

	fmt.Println(ans)

}
