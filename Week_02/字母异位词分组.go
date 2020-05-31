package main

/*
49. 字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{
			strs,
		}
	}
	chars := make([]int, 26)
	resMap := make(map[string][]string)
	res := [][]string{}
	for i := 0; i < len(strs); i++ {
		for k := 0; k < len(chars); k++ {
			chars[k] = 0
		}
		str := strs[i]
		for j := 0; j < len(str); j++ {
			chars[str[j]-'a']++
		}
		key := []byte{}
		for l := 0; l < len(chars); l++ {
			key = append(key, byte(chars[l]+'0'))
		}
		resMap[string(key)] = append(resMap[string(key)], str)
	}
	for _, v := range resMap {
		res = append(res, v)
	}
	return res
}
