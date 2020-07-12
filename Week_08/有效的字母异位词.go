package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charS := make([]int, 26)
	charT := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charS[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		charT[t[i]-'a']++
	}
	for i := 0; i < len(charT); i++ {
		if charT[i] != charS[i] {
			return false
		}
	}
	return true
}
