package day05

import (
	"strconv"
	"strings"
)

type Rule struct {
	first  int
	second int
}

func parseRule(input string) Rule {
	var rulestr []string = strings.Split(input, "|")
	first, err := strconv.Atoi(rulestr[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(rulestr[1])
	if err != nil {
		panic(err)
	}

	return Rule{first: first, second: second}
}

func parsePageNumbersUpdates(input string) (result []int) {
	var intstr []string = strings.Split(input, ",")
	for _, v := range intstr {
		parsedint, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result = append(result, parsedint)
	}
	return result
}

func isOrderGood(rules []Rule, input []int) bool {
	for _, rule := range rules {
		var foundRule []int
		for _, inp := range input {
			if rule.first == inp || rule.second == inp {
				foundRule = append(foundRule, inp)
			}
		}
		if len(foundRule) < 2 {
			continue
		}
		if foundRule[0] != rule.first && foundRule[1] != rule.second {
			// Found broken rule, second is before first
			return false
		}
	}
	return true
}
