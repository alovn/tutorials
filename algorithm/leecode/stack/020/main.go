package main

import "fmt"

func main() {
	s := "([)"
	s = "()[]{}"
	s = "(]"
	s = "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []rune
	m := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, c := range s {
		if _, ok := m[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			latest := stack[len(stack)-1]
			p, ok := m[latest]
			if !ok {
				return false
			}
			if c != p {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
