package main

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	// dp[i][j] 代表word1 1-i 变换到 word2 1-j 需要几步
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	// 当word1为空 变到 word2 每一步需要添加1个字母 所以+1
	for i := 1; i <= n2; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	// 当word2为空 word1 变到word2 每一步需要删除一个字母 所以操作次数+1
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				// 从1开始 当word1 第i 等于word2 第j
				// 说明只需要知道i前面 和j前面的单词经过多少就行
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 求dp[i][j] 有三种变换：
				// 以 word1 为 "horse"，word2 为 "ros"，且 dp[5][3] 为例，即要将 word1的前 5 个字符转换为 word2的前 3 个字符，也就是将 horse 转换为 ros，因此有：
				//(1) dp[i-1][j-1]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 2 个字符 ro，然后将第五个字符 word1[4]（因为下标基数以 0 开始） 由 e 替换为 s（即替换为 word2 的第三个字符，word2[2]）
				//(2) dp[i][j-1]，即先将 word1 的前 5 个字符 horse 转换为 word2 的前 2 个字符 ro，然后在末尾补充一个 s，即插入操作
				//(3) dp[i-1][j]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 3 个字符 ros，然后删除 word1 的第 5 个字符
				dp[i][j] = minTrans(dp[i-1][j-1], minTrans(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[n1][n2]
}

func minTrans(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
