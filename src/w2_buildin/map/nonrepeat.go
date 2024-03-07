package main

// 寻找最长不含有重复字符串的子串（不支持中文）
func lenOfNonRepeatSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLen := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLen
}

// 寻找最长不含有重复字符串的子串（支持中文）
func lenOfNonRepeatSubStrGlobal(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLen := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLen
}
