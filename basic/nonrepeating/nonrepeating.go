package main

// abcadefg
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccureed := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccureed[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccureed[ch] = i
	}
	return maxLength
}
