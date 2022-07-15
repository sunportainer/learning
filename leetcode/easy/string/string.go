package string

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	chars := []rune(strings.ToLower(s))
	begin, end := 0, len(s)-1
	for begin < end {
		for begin < end && !(unicode.IsDigit(chars[begin]) || unicode.IsLetter(chars[begin])) {
			begin = begin + 1
		}
		for begin < end && !(unicode.IsDigit(chars[end]) || unicode.IsLetter(chars[end])) {
			end = end - 1
		}

		if chars[begin] != chars[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func TestisPalindrome() {
	str1 := "A man, a plan, a canal: Panama"
	ret := isPalindrome(str1)
	fmt.Println(ret)

	str2 := "\"race a car\""
	ret = isPalindrome(str2)
	fmt.Println(ret)

	str3 := ".,"
	ret = isPalindrome(str3)
	fmt.Println(ret)
}

/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
s = "leetcode"  result = 0
s = "loveleetcode" result = 2
*/
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	hash := make(map[byte]int, 26)
	chars := []byte(s)
	for index, val := range chars {
		if _, ok := hash[val]; ok {
			hash[val] = -2
		} else {
			hash[val] = index
		}
	}
	min := len(s)
	find := false
	for _, v := range hash {
		if v != -2 {
			if v <= min {
				min = v
				find = true
			}
		}
	}
	if !find {
		min = -1
	}
	return min

}

func TestfirstUniqChar() {
	s := "leetcode"
	ret := firstUniqChar(s)
	fmt.Println(ret)
	s = "loveleetcode"
	ret = firstUniqChar(s)
	fmt.Println(ret)
}
