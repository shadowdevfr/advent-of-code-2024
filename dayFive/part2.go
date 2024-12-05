package dayfive

import "sort"

func fixOrder(rules []Rule, input []int) []int {
	sort.SliceStable(input, func(i, j int) bool {
		for _, rule := range rules {
			if (input[i] == rule.first && input[j] == rule.second) || (input[i] == rule.second && input[j] == rule.first) {
				return input[i] == rule.first
			}
		}
		return input[i] < input[j]
	})
	return input
}
