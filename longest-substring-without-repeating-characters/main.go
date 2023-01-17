// Given a string s, find the length of the longest substring without repeating characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// Constraints:
// 0 <= s.length <= 5 * 10^4
// s consists of English letters, digits, symbols and spaces.
// See https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func contains(slice string, elem string) bool {
	for i := 0; i < len(slice); i++ {
		if string(slice[i]) == elem {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	longestSubstring := ""
	cursorPos := 0

	for cursorPos < len(s) {
		i := cursorPos
		curLongestSubstring := ""
		for i < len(s) {
			if contains(curLongestSubstring, string(s[i])) {
				break
			}
			curLongestSubstring += string(s[i])
			i++
		}
		if len(curLongestSubstring) > len(longestSubstring) {
			longestSubstring = curLongestSubstring
		}
		if len(longestSubstring) > len(s)-cursorPos {
			break
		}
		cursorPos++
	}
	/* Uncomment this to print the actual longest substring without repeating characters
	fmt.Printf("longest substring: \"%s\"\r\n", longestSubstring)
	*/
	return len(longestSubstring)
}

func main() {
	// Testing "Case 1"
	var s1 string = "abcabcbb"
	fmt.Printf("s=\"%s\"\r\n", s1)
	fmt.Printf("length=%d\r\n\r\n", lengthOfLongestSubstring(s1))

	// Testing "Case 2"
	var s2 string = "bbbbb"
	fmt.Printf("s=\"%s\"\r\n", s2)
	fmt.Printf("length=%d\r\n\r\n", lengthOfLongestSubstring(s2))

	// Testing "Case 3"
	var s3 string = "pwwkew"
	fmt.Printf("s=\"%s\"\r\n", s3)
	fmt.Printf("length=%d\r\n\r\n", lengthOfLongestSubstring(s3))

	// TODO: add the code to wait for user's input via the terminal window
}
