package main

// dp
func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[:1]
		}
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				dp[i][i+1] = true
			}
		}
	}
	keyA, keyB := 0, 0
	res := 1
	// 固定尾部j，然后从头往后扫i=[0,j)
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if dp[i][j] {
				if j-i+1 > res {
					keyA = i
					keyB = j
					res = j - i + 1
				}
			} else {
				if dp[i+1][j-1] {
					if s[i] == s[j] {
						dp[i][j] = true
						if j-i+1 > res {
							keyA = i
							keyB = j
							res = j - i + 1
						}
					}
				}
			}
		}
	}
	return s[keyA : keyB+1]
}

// 中心拓展
func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[:len(s)-1]
		}
	}
	res := ""
	maxL := 0
	for i := 0; i < len(s)-1; i++ {
		oddStr := centerStr(s, i, i)
		evenStr := centerStr(s, i, i+1)
		if len(oddStr) > maxL {
			maxL = len(oddStr)
			res = oddStr
		}
		if len(evenStr) > maxL {
			maxL = len(evenStr)
			res = evenStr
		}
	}
	return res
}

func centerStr(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return s[l+1 : r]
}
