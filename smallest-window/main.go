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
	rightResult := RightBackSubstring(main, substring)
	leftResult := LeftSubstring(main, substring)

	finalLeft := RightBackSubstring(leftResult, substring)
	finalRight := LeftSubstring(rightResult, substring)

	if len(finalLeft) < len(finalRight) {
		return finalLeft
	}
	return finalRight
}

func RightBackSubstring(main, sub string) (resultStr string) {
	
	fmt.Println(main)
	resultStr = main[:len(main)-1]
	
	rChecker := PopulateCharChecker(resultStr)

 	if !doesContainAllChars(sub, rChecker) {
		return main
	} 

	return RightBackSubstring(resultStr, sub)
}	

func LeftSubstring(main, sub string) (resultStr string) {
	fmt.Println(main)
	resultStr = main[1:]
	
	lChecker := PopulateCharChecker(resultStr)

 	if !doesContainAllChars(sub, lChecker) {
		return main
	}
	return LeftSubstring(resultStr, sub)

}

func main() {
	
	input :=  "ADOCBECODEBANCZ"
	substring := "ABC"

	fmt.Println("This is the challenge start!")
	resultStr :=  GetSmallestWindow(input, substring)
	fmt.Println("final result = ", resultStr)


}