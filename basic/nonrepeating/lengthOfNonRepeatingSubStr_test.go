package main

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct {
		str       string
		maxLength int
	}{
		//normal
		{"abcd", 4},
		{"abcdea", 5},

		//
		{"aaa", 1},
		{"", 0},
		{"aaaaaaaab", 2},
		{"中文mooc", 4},
	}
	for _, v := range tests {
		if actual := lengthOfNonRepeatingSubStr(v.str); actual != v.maxLength {
			t.Errorf("Error occupied when test:%s, got answer:%d, expected: %d \n", v.str, actual, v.maxLength)
		}

	}
}

func BenchmarkSubStr(b *testing.B) {
	str, maxlength := "中国Mooc", 4
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(str)
		if actual != maxlength {
			b.Errorf("got %d for input %s; expected: %d", actual, str, maxlength)
		}
	}
}
