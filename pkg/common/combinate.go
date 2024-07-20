package common

import "strings"

func Combinate(words []string, current []string, result *[]string) {
	if len(current) > 0 {
		*result = append(*result, strings.Join(current, ","))
	}
	for i := 0; i < len(words); i++ {
		next := append([]string{}, current...)
		next = append(next, words[i])
		remaining := append([]string{}, words[i+1:]...)
		Combinate(remaining, next, result)
	}
}
