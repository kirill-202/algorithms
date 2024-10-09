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


func GetSmallestWindow(main, substring string) string {
	checker := PopulateCharChecker(main)
	if !doesContainAllChars(substring, checker) {
		return ""
	}

	return  SlidingCheck(main, substring, len(substring))
}

func SlidingCheck(main, sub string, length int) (resultStr string) {

	
	//shift substring chunks up to the current length-1
	for i:=0;i<len(main)-length; i++ {
		chunk := main[i:length+i]
		tempChecker := PopulateCharChecker(chunk)
		if doesContainAllChars(sub, tempChecker) {
			return chunk
		}

	}
	length++

	return SlidingCheck(main, sub, length)
}	

func main() {
	
	input :=  "CBAAAADOCBEAECODEBANCZ"
	substring := "ABC"

	fmt.Println("This is the challenge start!")
	resultStr :=  GetSmallestWindow(input, substring)
	fmt.Println("final result = ", resultStr)


}