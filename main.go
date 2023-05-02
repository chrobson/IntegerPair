package main

import (
	"fmt"
)

func findPair(arr []int, sum int) [][]int {
	n := len(arr)
	pairs := [][]int{}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		complement := sum - arr[i]
		if _, ok := m[complement]; ok {
			pairs = append(pairs, []int{arr[i], complement})
		}
		m[arr[i]] = i
	}
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1]) {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
	return pairs
}

func main() {
	nums := []int{8, 7, 2, 5, 3, 1}
	target := 10
	pairs := findPair(nums, target)
	fmt.Println(pairs)
}
