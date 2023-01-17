// You are given a string s and an array of strings words. All the strings of words are of the same length.
// A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.
// Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.
// Example 1:
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Output: [0,9]
// Example 2:
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Output: []
// Example 3:
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// Output: [6,9,12]
// See https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package main

import "fmt"

func findSubstring(s string, words []string) []int {
	numOfWords := len(words)
	if numOfWords < 1 {
		return []int{}
	}
	wordLength := len(words[0])
	sLength := len(s)
	if sLength < wordLength*numOfWords {
		return []int{}
	}
	cursorPos := 0
	var res []int = []int{}
	for cursorPos < sLength {
		tmpWords := make([]string, numOfWords)
		copy(tmpWords, words)
		i := cursorPos
		for i < sLength {
			isMatched := false
			for j := 0; j < len(tmpWords); j++ {
				if len(s) < i+wordLength {
					break
				}
				if tmpWords[j] == string(s[i:i+wordLength]) {
					isMatched = true
					tmpWords = append(tmpWords[:j], tmpWords[j+1:]...)
					break
				}
			}
			if !isMatched {
				break
			}
			i += wordLength
			if len(tmpWords) == 0 {
				res = append(res, cursorPos)
				break
			}
		}
		cursorPos++
		if sLength-cursorPos < wordLength*numOfWords {
			break
		}
	}
	return res
}

func main() {
	s1 := "barfoothefoobarman"
	fmt.Printf("s=\"%s\"\r\n", s1)
	words1 := []string{"foo", "bar"}
	fmt.Println("words: ", words1)
	fmt.Println("positions of the substring with concatenation of all words: ", findSubstring(s1, words1))

	s2 := "wordgoodgoodgoodbestword"
	fmt.Printf("s=\"%s\"\r\n", s2)
	words2 := []string{"word", "good", "best", "word"}
	fmt.Println("words: ", words2)
	fmt.Println("positions of the substring with concatenation of all words: ", findSubstring(s2, words2))

	s3 := "barfoofoobarthefoobarman"
	fmt.Printf("s=\"%s\"\r\n", s1)
	words3 := []string{"bar", "foo", "the"}
	fmt.Println("words: ", words3)
	fmt.Println("positions of the substring with concatenation of all words: ", findSubstring(s3, words3))

	// TODO: add the code to wait for user's input via the terminal window
}
