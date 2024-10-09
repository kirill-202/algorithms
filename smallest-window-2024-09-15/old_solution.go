package main

import (
  "fmt"

)



func PopulateChecker(inputString string) map[rune]bool {
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
  checker := PopulateChecker(main)
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
  
  rChecker := PopulateChecker(resultStr)

   if !doesContainAllChars(sub, rChecker) {
    return main
  } 

  return RightBackSubstring(resultStr, sub)
}  

func LeftSubstring(main, sub string) (resultStr string) {
  fmt.Println(main)
  resultStr = main[1:]
  
  lChecker := PopulateChecker(resultStr)

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