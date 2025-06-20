package a17countthenumberofvowelstringsinrange2586

import (
	"fmt"
	"strings"
)

func vowelStrings(words []string, left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			fmt.Println(words[i])
			cnt++
		}
	}
	return cnt
}

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}

func vowelStrings2(words []string, left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		if strings.Contains("aeiou", words[i][:1]) && strings.Contains("aeiou", words[i][(len(words[i])-1):]) {
			cnt++
		}
	}
	return cnt
}

func vowelStrings3(words []string, left int, right int) int {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	cnt := 0
	for i := left; i <= right; i++ {
		if _, ok := vowels[words[i][0]]; ok {
			if _, ok := vowels[words[i][len(words[i])-1]]; ok {
				cnt++
			}
		}
	}
	return cnt
}
