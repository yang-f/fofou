// This code is under BSD license. See license-bsd.txt
package main

import "strings"

func isSp(c rune) bool {
	return c == ' '
}

func isNewline(s string) bool {
	return 1 == len(s) && s[0] == '\n'
}

func isNewlineChar(c rune) bool {
	return c == '\n'
}

func endsSendence(s string) bool {
	n := len(s)
	if 0 == n {
		return false
	}
	c := s[n-1]
	if c == '.' || c == '?' || c == '\n' {
		return true
	}
	return false
}

// TODO: this is a bit clumsy. Would be much faster (and probably cleaner) to
// go over string char-by-char
// TODO: only do it if detects high CAPS rate
func UnCaps(s string) string {
	parts := strings.FieldsFunc(s, isSp)
	n := len(parts)
	res := make([]string, n, n)
	sentenceStart := true
	for i := 0; i < n; i++ {
		s := parts[i]
		if isNewline(s) {
			res[i] = s
			sentenceStart = true
			continue
		}
		s2 := strings.ToLower(s)
		if sentenceStart {
			res[i] = strings.Title(s2)
		} else {
			res[i] = s2
		}
		sentenceStart = endsSendence(s)
	}
	s = strings.Join(res, " ")
	return s
	/*
		parts = strings.FieldsFunc(s, isNewlineChar)
		n = len(parts)
		res = make([]string, n, n)
		for i := 0; i < n; i++ {
			res[i] = strings.Title(res[i])
		}
		return strings.Join(res, "\n")
	*/
}
