package main

import "fmt"

var matched bool = false

func main() {
	match("abbbbc", "a*")
	fmt.Print(matched)
	fmt.Println()
}

func match(s string, pattern string) bool {
	rmatch(0, 0, len(s), len(pattern), s, pattern)
	return matched
}

func rmatch(si int, pi int, slen int, plen int, s string, pattern string) {
	if matched {
		return
	}
	if pi == plen {
		if si == slen {
			matched = true
		}
		return
	}

	if pattern[pi] == '*' {
		for k := 0; k <= slen-si; k++ {
			rmatch(si+k, pi+1, slen, plen, s, pattern)
		}
	} else if pattern[pi] == '?' {
		rmatch(si, pi+1, slen, plen, s, pattern)
		rmatch(si+1, pi+1, slen, plen, s, pattern)
	} else if si < slen && pattern[pi] == s[si] {
		rmatch(si+1, pi+1, slen, plen, s, pattern)
	}
}
