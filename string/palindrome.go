package main

import "strings"

func isPalindrome(s string) bool {
	s = preProcess(s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func preProcess(s string) string {
	var str strings.Builder
	for _, c := range s {
		if c >= 'a' && c <= 'z' { // lower case
			str.WriteRune(c - ('a' - 'A'))
		} else if c >= 'A' && c <= 'Z' { // upper case
			str.WriteRune(c)
		} else if c >= '0' && c <= '9' { // digit
			str.WriteRune(c)
		}
	}
	return str.String()
}
