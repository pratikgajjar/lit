package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	right := []int{}

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
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	ans := 0

	for i, _ := range left {
		ans += abs(left[i] - right[i])
	}

	fmt.Println(ans)

}
