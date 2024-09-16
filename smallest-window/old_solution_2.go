package main

import (
	"fmt"
)

/*
Problem Description:
Given two strings S and T, find the minimum window in S which will contain all the characters in T
 in complexity better than O(n^2).

Input:
A string S of length n (1 ≤ n ≤ 10^6), consisting of lowercase and uppercase letters.
A string T of length m (1 ≤ m ≤ 10^4), consisting of lowercase and uppercase letters.
Output:
Return the minimum window in S that contains all the characters from T.
If there is no such window, return an empty string. If there are multiple windows of the same length, return the first one found.


*/


func PopulateCharChecker(inputString string) map[rune]bool {
	runeChecker := make(map[rune]bool)
    for _, char := range inputString {  

        _, exist := runeChecker[char]
        if exist {
            continue
        }
        runeChecker[char] = true
    }
	return runeChecker
}

func doesContainAllChars(substring string, runeChecker map[rune]bool) bool {

	for  _, char := range substring {
		_, exist := runeChecker[char]
		if !exist {
			return false
		}
	}
	return true
}

func chunkString(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		return nil
	}
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}




func GetSmallestWindow(main, substring string) string {
	checker := PopulateCharChecker(main)
	if !doesContainAllChars(substring, checker) {
		return ""
	}

	return  GetSubstring(main, substring, len(substring))
}

func GetSubstring(main, sub string, length int) (resultStr string) {

	if length == len(main) {
		return main
	}

	testSubstrings := chunkString(main, length)
	
	//shift substring chunks up to the current length-1
	iters := length-1

	for i:=1; iters>0; i++ {
		iters--
		additonalStep := chunkString(main[i:], length)
		testSubstrings = append(testSubstrings, additonalStep...)
	}
	fmt.Println(testSubstrings)

	for _, chunk := range testSubstrings {
		tempChecker := PopulateCharChecker(chunk)
		if doesContainAllChars(sub, tempChecker) {
			return chunk
		} 
	}
	
	length++

	return GetSubstring(main, sub, length)
}	


func main() {
	
	input :=  "AAAADOCBEAECODEBANCZ"
	substring := "ABC"

	fmt.Println("This is the challenge start!")
	resultStr :=  GetSmallestWindow(input, substring)
	fmt.Println("final result = ", resultStr)


}